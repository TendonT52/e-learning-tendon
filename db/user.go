package db

import (
	"context"
	"errors"
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

func (u *userDB) InsertUser(user *core.User) error {
	userDoc := userDoc{
		ID:             primitive.NewObjectID(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
		Curricula:      HexIDToObjID(user.Curricula),
		UpdatedAt:      primitive.NewDateTimeFromTime(time.Now()),
	}
	user.ID = userDoc.ID.Hex()
	user.UpdatedAt = userDoc.UpdatedAt.Time()

	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err := u.collection.InsertOne(ctx, userDoc)
	if err != nil {
		return errs.InsertFailed.From(err)
	}
	return nil
}

func (u *userDB) InsertManyUser(users []core.User) error {
	userDocs := make([]interface{}, len(users))
	for i, user := range users {
		userDoc := userDoc{
			ID:             primitive.NewObjectID(),
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			Email:          user.Email,
			HashedPassword: user.HashedPassword,
			Role:           user.Role,
			Curricula:      HexIDToObjID(user.Curricula),
			UpdatedAt:      primitive.NewDateTimeFromTime(time.Now()),
		}
		users[i].ID = userDoc.ID.Hex()
		users[i].UpdatedAt = userDoc.UpdatedAt.Time()
		userDocs[i] = userDoc
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err := u.collection.InsertMany(ctx, userDocs)
	if err != nil {
		return errs.InsertFailed.From(err)
	}
	return nil
}

func (u *userDB) FindUserByEmail(email string) (core.User, error) {
	filter := bson.M{"email": email}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var userDoc userDoc
	err := u.collection.FindOne(ctx, filter).Decode(&userDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.User{}, errs.UserNotFound
		}
		return core.User{}, errs.FindFailed.From(err)
	}
	user := userDoc.toUser()
	return user, nil
}

func (u *userDB) FindUser(hexID string) (core.User, error) {
	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return core.User{}, errs.InvalidUserID.From(err)
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var userDoc userDoc
	err = u.collection.FindOne(ctx, filter).Decode(&userDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.User{}, errs.UserNotFound
		}
		return core.User{}, errs.FindFailed.From(err)
	}
	user := userDoc.toUser()
	return user, nil
}

func (u *userDB) FindManyUser(hexIDs []string) ([]core.User, error) {
	objIDs := make([]primitive.ObjectID, len(hexIDs))
	for i, hexID := range hexIDs {
		objID, err := primitive.ObjectIDFromHex(hexID)
		if err != nil {
			return nil, errs.InvalidUserID.From(errors.New(
				"invalid user id: " + hexID,
			))
		}
		objIDs[i] = objID
	}
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	cursor, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.FindFailed.From(err)
	}
	var userDocs []userDoc
	if err = cursor.All(ctx, &userDocs); err != nil {
		return nil, errs.FindFailed.From(err)
	}
	users := make([]core.User, len(userDocs))
	for i, userDoc := range userDocs {
		users[i] = userDoc.toUser()
	}
	return users, nil
}

func (u *userDB) UpdateUser(user *core.User) error {
	userObjID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return errs.InvalidUserID.From(err)
	}
	user.UpdatedAt = time.Now()
	update := bson.M{"$set": bson.M{
		"firstName":      user.FirstName,
		"lastName":       user.LastName,
		"email":          user.Email,
		"hashedPassword": user.HashedPassword,
		"role":           user.Role,
		"curricula":      HexIDToObjID(user.Curricula),
		"updatedAt":      primitive.NewDateTimeFromTime(user.UpdatedAt),
	}}
	result, err := u.collection.UpdateByID(context.Background(), userObjID, update)
	if err != nil {
		return errs.UpdateFailed.From(err)
	}
	if result.MatchedCount == 0 {
		return errs.UserNotFound
	}
	return nil
}

func (u *userDB) DeleteUser(hexID string) error {
	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return errs.InvalidUserID.From(err)
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.UserNotFound
	}
	return nil
}

func (u *userDB) DeleteManyUser(hexIDs []string) error {
	ObjIDs := HexIDToObjID(hexIDs)
	filter := bson.M{"_id": bson.M{"$in": ObjIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	_, err := u.collection.DeleteMany(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return errs.DeleteFailed.From(err)
	}
	return nil
}

func (u *userDB) Clear() {
	u.collection.DeleteMany(context.Background(), bson.M{})
}
