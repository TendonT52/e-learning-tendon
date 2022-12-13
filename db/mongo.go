package db

import (
	"context"
	"log"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var config MongoConfig

type MongoConfig struct {
	CreateTimeOut time.Duration
	FindTimeout   time.Duration
	UpdateTimeout time.Duration
	DeleteTimeout time.Duration
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

func ArrayStringToArrayObjectId(arr []string) ([]primitive.ObjectID, error) {
	obj := make([]primitive.ObjectID, len(arr))
	for i, v := range arr {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, errs.ErrWrongFormat.From(err)
		}
		obj[i] = id
	}
	return obj, nil
}

func ArrayObjectIdToArrayString(arr []primitive.ObjectID) []string {
	ln := make([]string, len(arr))
	for i, v := range arr {
		ln[i] = v.Hex()
	}
	return ln
}
