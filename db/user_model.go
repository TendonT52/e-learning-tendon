package db

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userDoc struct {
	ID             primitive.ObjectID   `bson:"_id"`
	FirstName      string               `bson:"firstName"`
	LastName       string               `bson:"lastName"`
	Email          string               `bson:"email"`
	HashedPassword string               `bson:"hashed_password"`
	Role           string               `bson:"role"`
	Curricula      []primitive.ObjectID `bson:"curricula"`
	UpdatedAt      primitive.DateTime   `bson:"updated_at"`
}

func (u *userDoc) toUser() core.User {
	return core.User{
		ID:             u.ID.Hex(),
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		Role:           u.Role,
		Courses:        ObjIDToHexID(u.Curricula),
		UpdatedAt:      u.UpdatedAt.Time(),
	}
}
