package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lighty0410/ekadashi-server/pkg/mongo"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	loginRequest
}

// handleRegistration registers user in the system.
func (s *EkadashiServer) handleRegistration(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("can not decode the request: %v", err))
		return
	}
	hashedPassword, err := generateHash(req.Password)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password: %v", err))
		return
	}
	err = s.db.AddUser(&mongo.User{
		Name: req.Username,
		Hash: hashedPassword,
	})
	if err != nil {
		jsonError(w, http.StatusInternalServerError, fmt.Errorf("could not add user: %v", err))
		return
	}
	jsonResponse(w, http.StatusOK, nil)
}

func (s *EkadashiServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("can not decode the request: %v", err))
		return
	}
	user, err := s.db.ReadUser(req.Username)
	if err == mongo.ErrUserNotFound {
		jsonError(w, http.StatusUnauthorized, fmt.Errorf("incorrect username or password: %v", err))
		return
	}
	if err != nil {
		jsonError(w, http.StatusInternalServerError, fmt.Errorf("an error occurred in mongoDB: %v", err))
		return
	}
	err = compareHash(user.Hash, []byte(req.Password))
	if err != nil {
		jsonError(w, http.StatusUnauthorized, fmt.Errorf("incorrect username or password"))
		return
	}
	jsonResponse(w, http.StatusOK, nil)
}
