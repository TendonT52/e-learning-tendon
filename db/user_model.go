package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userDoc struct {
	Id               primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName        string               `bson:"firstName"`
	LastName         string               `bson:"lastName"`
	Email            string               `bson:"email"`
	HashPassword     string               `bson:"password"`
	Role             string               `bson:"role"`
	UpdatedAt        primitive.DateTime   `bson:"updated_at"`
	EnrollCurriculum []primitive.ObjectID `bson:"enroll"`
}
