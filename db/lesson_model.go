package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type lessonDoc struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Access      string               `bson:"access"`
	CreateBy    primitive.ObjectID   `bson:"create_by"`
	UpdatedAt   primitive.DateTime   `bson:"update_at"`
	Node        []primitive.ObjectID `bson:"node"`
	NextLesson  []primitive.ObjectID `bson:"next"`
	PrevLesson  []primitive.ObjectID `bson:"prev"`
}
