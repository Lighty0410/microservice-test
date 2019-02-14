package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Service is used to interact with ekadashi storage.
type Service struct {
	db *mongo.Database
}

// NewService attempts to connect to MongoDB at localhost and
// if connection succeeds it returns Service ready to use.
func NewService(connectionURL string) (*Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, connectionURL)
	if err != nil {
		return nil, fmt.Errorf("could not dial mongo: %v", err)
	}
	db := client.Database("ekadashi")
	s := &Service{
		db: db,
	}
	err = s.CreateIndex()
	if err != nil{
		return s, fmt.Errorf("cannot create an index")
	}
	return s, nil
}

// CreateIndex creates an index for session collection
func (s *Service) CreateIndex() error {
	var opt options.IndexOptions
	opt.SetExpireAfterSeconds(60*5)
	model := mongo.IndexModel{
		Options: &opt,
	}
	c := s.db.Collection("session")
	_, err := c.Indexes().CreateOne(context.Background(), model)
	if err != nil {
		return fmt.Errorf("cannot create an index")
	}
	return nil
}
