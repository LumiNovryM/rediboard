package db

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("NO MATCHING RECORD FOUND IN REDIS DATABASE")
	Ctx = context.TODO()
)

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 15,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err;
	}

	return &Database{
		Client: client,
	}, nil
}