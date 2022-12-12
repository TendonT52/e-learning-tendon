package db

import (
	"context"
	"log"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserDBInstance *userDB

type userDB struct {
	collection *mongo.Collection
}

func NewUserDB(userCollectionName string) {
	UserDBInstance = &userDB{
		db.Collection(userCollectionName),
	}
}

func (u *userDB) InsertUser(firstName, lastName, email, hashPassword, role string) (core.User, error) {
	doc := userDoc{
		Id:           primitive.NewObjectID(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		HashPassword: hashPassword,
		Role:         role,
		UpdatedAt:    primitive.NewDateTimeFromTime(time.Now()),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err := u.collection.InsertOne(ctx, doc)
	if err != nil {
		return core.User{}, errs.ErrDatabase.From(err)
	}
	user := core.User{
		ID:           doc.Id.Hex(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		HashPassword: hashPassword,
		Role:         role,
		UpdatedAt:    doc.UpdatedAt.Time(),
	}
	return user, nil
}

func (u *userDB) GetUserByEmail(email string) (core.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := userDoc{}
	err := u.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.User{}, errs.ErrNotFound.From(err)
		}
		return core.User{}, errs.ErrDatabase.From(err)
	}
	user := core.User{
		ID:           doc.Id.Hex(),
		FirstName:    doc.FirstName,
		LastName:     doc.LastName,
		Email:        doc.Email,
		HashPassword: doc.HashPassword,
		Role:         doc.Role,
		UpdatedAt:    doc.UpdatedAt.Time(),
	}
	return user, nil
}

func (u *userDB) GetUserById(id string) (core.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.User{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := userDoc{}
	err = u.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.User{}, errs.ErrNotFound.From(err)
		}
		return core.User{}, errs.ErrDatabase.From(err)
	}
	user := core.User{
		ID:           doc.Id.Hex(),
		FirstName:    doc.FirstName,
		LastName:     doc.LastName,
		Email:        doc.Email,
		HashPassword: doc.HashPassword,
		Role:         doc.Role,
		UpdatedAt:    doc.UpdatedAt.Time(),
	}
	return user, nil
}

func (u *userDB) CleanUp() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	result, err := u.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up user collection, %v", err)
	}
	return int(result.DeletedCount)
}
