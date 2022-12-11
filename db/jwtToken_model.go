package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type jwtDoc struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Exp primitive.DateTime `bson:"exp"`
}
