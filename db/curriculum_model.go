package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type curriculumDoc struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"curriculum_name"`
	Description string               `bson:"description"`
	Access      string               `bson:"access"`
	CreateBy    primitive.ObjectID   `bson:"create_by"`
	UpdatedAt   primitive.DateTime   `bson:"updated_at"`
	Lessons     []primitive.ObjectID `bson:"lessons"`
}
