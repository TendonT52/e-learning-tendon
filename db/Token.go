package db

import (
	"context"
	"log"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var JwtDBInstance *jwtDB

type jwtDB struct {
	collection *mongo.Collection
}

func NewJwtDB(jsontokenCollectionName string) {
	JwtDBInstance = &jwtDB{
		db.Collection(jsontokenCollectionName),
	}
}

func (jw *jwtDB) InsertJwt(exp time.Time) (string, error) {
	doc := jwtDoc{
		Id:  primitive.NewObjectID(),
		Exp: primitive.NewDateTimeFromTime(exp),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	result, err := jw.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", errs.InsertFailed.From(err)
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (jw *jwtDB) CheckJwt(hexId string) error {
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidToken.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	doc := jwtDoc{}
	err = jw.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.TokenNotfound.From(err)
		}
		return errs.FindFailed.From(err)
	}
	return nil
}

func (jw *jwtDB) DeleteJwt(hexId string) error {
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidToken.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := jw.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.TokenNotfound
	}
	return nil
}

func (jw *jwtDB) Clear() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	result, err := jw.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up jwt collection, %v", err)
	}
	return int(result.DeletedCount)
}
