package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type nodeDoc struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `bson:"type"`
	Data      string             `bson:"data"`
	CreateBy  primitive.ObjectID `bson:"create_by"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
