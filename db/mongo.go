package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var config MongoConfig

type MongoConfig struct {
	InsertTimeOut time.Duration
	FindTimeOut   time.Duration
	UpdateTimeOut time.Duration
	DeleteTimeOut time.Duration
}

func NewClient(connection string, mongoconfig MongoConfig) {
	config = mongoconfig
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

func ObjIDToHexID(arr []primitive.ObjectID) []string {
	ln := make([]string, len(arr))
	for i, v := range arr {
		ln[i] = v.Hex()
	}
	return ln
}

func HexIDToObjID(arr []string) []primitive.ObjectID {
	ln := make([]primitive.ObjectID, len(arr))
	for i, v := range arr {
		ln[i], _ = primitive.ObjectIDFromHex(v)
	}
	return ln
}
