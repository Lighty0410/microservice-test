package ekadashihttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Lighty0410/ekadashi-server/pkg/server/controller"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	loginRequest
}

var validPassword = regexp.MustCompile(`^[a-zA-Z1-9=]+$`).MatchString
var validUsername = regexp.MustCompile(`^[a-zA-Z1-9]+$`).MatchString

// handleRegistration registers user in the system.
func (s *EkadashiServer) handleRegistration(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("can not decode the request: %v", err))
		return
	}
	err = req.validateRequest()
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	response, err := s.controller.RegisterUser(controller.User{Username: req.Username, Password: req.Password})
	switch response {
	case controller.StatusOK:
		jsonResponse(w, http.StatusOK, nil)
		return
	case controller.StatusInternalServerError:
		jsonError(w, http.StatusInternalServerError, err)
		return
	case controller.StatusConflict:
		jsonError(w, http.StatusConflict, err)
		return
	case controller.StatusBadRequest:
		jsonError(w, http.StatusBadRequest, err)
		return
	}
}

// handleLogin retrieves information about login request.
// If login succeed it assigns cookie to user.
func (s *EkadashiServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("can not decode the request: %v", err))
		return
	}
	err = req.validateRequest()
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	response, cookieValues, err := s.controller.LoginUser(controller.User{Username: req.Username, Password: req.Password})
	switch response {
	case controller.StatusUnauthorized:
		jsonError(w, http.StatusUnauthorized, err)
		return
	case controller.StatusInternalServerError:
		jsonError(w, http.StatusInternalServerError, err)
		return
	case controller.StatusConflict:
		jsonError(w, http.StatusConflict, err)
		return
	case controller.StatusBadRequest:
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	cookie := http.Cookie{
		Name:  cookieValues.Name,
		Value: cookieValues.SessionHash,
	}
	http.SetCookie(w, &cookie)
	jsonResponse(w, http.StatusOK, nil)
}
