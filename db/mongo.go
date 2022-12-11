package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

var createTimeOut time.Duration = time.Second * 5
var findTimeout time.Duration = time.Second * 5
var updateTimeout time.Duration = time.Second * 5
var deleteTimeout time.Duration = time.Second * 5

func NewClient(connection string) {
	log.Println("Connecting to mongo...")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connection))
	if err != nil {
		log.Printf("Connection String, %v \n", connection)
		log.Fatalf("Error while connect to mongo, %v", err)
	}
	err = pingMongo()
	if err != nil {
		log.Fatalf("Error while ping to mongo, %v", err)
	}
	log.Println("Connect mongo completed")
}

func NewDB(dbName string) {
	db = client.Database(dbName)
}

func DisconnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Error while disconnect mongo, %v", err)
	}
	log.Println("Disconnect mongo completed")
}

func pingMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := client.Ping(ctx, nil)
	return err
}

