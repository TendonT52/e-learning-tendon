package db

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type lessonDoc struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Access      string               `bson:"access"`
	CreateBy    primitive.ObjectID   `bson:"create_by"`
	UpdatedAt   primitive.DateTime   `bson:"update_at"`
	Nodes       []primitive.ObjectID `bson:"nodes"`
	NextLessons []primitive.ObjectID `bson:"next_lessons"`
	PrevLessons []primitive.ObjectID `bson:"prev_lessons"`
}

func (l lessonDoc) toLesson() core.Lesson {
	return core.Lesson{
		ID:          l.ID.Hex(),
		Name:        l.Name,
		Description: l.Description,
		Access:      l.Access,
		CreateBy:    l.CreateBy.Hex(),
		UpdatedAt:   l.UpdatedAt.Time(),
		Nodes:       ObjIDToHexID(l.Nodes),
		NextLessons: ObjIDToHexID(l.NextLessons),
		PrevLessons: ObjIDToHexID(l.PrevLessons),
	}
}
