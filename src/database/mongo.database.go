package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client  *mongo.Client
	Users   *mongo.Collection
	Resumes *mongo.Collection
	About   *mongo.Collection
)

func Connect(uri, database string) error {
	clientOptions := options.Client().ApplyURI(uri)

	localClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	client = localClient

	Users = client.Database(database).Collection("users")
	Resumes = client.Database(database).Collection("resumes")
	About = client.Database(database).Collection("about")

	err = localClient.Database(database).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	return err
}

func Close() error {
	return client.Disconnect(context.Background())
}
