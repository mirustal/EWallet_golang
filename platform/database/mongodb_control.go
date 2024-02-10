package database

import (
	"context"
	"ewallet/pkg/configs"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db_mongo *mongo.Database

func NewClient(ctx context.Context) (db *mongo.Database, err error) {
	cfg := configs.GetConfig()
	host, port, database := cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Database
	mongoDBURI := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOptions := options.Client().ApplyURI(mongoDBURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect MongoDB")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB")
	}

	db_mongo = client.Database(database)

	return db_mongo, nil

}