package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Config struct {
	URI     string
	Timeout int
}

func NewDriver(config Config) (*mongo.Client, error) {
	if len(config.URI) == 0 {
		return nil, errors.New("invalid_mongo_uri")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}
	log.Println("db successfully connected")

	//defer client.Disconnect(ctx)

	return client, nil
}
