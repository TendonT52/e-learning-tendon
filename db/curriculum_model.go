package db

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type curriculumDoc struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Access      string               `bson:"access"`
	CreateBy    primitive.ObjectID   `bson:"create_by"`
	UpdatedAt   primitive.DateTime   `bson:"updated_at"`
	Lessons     []primitive.ObjectID `bson:"lessons"`
}

func (c curriculumDoc) toCurriculum() core.Curriculum {
	return core.Curriculum{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Description: c.Description,
		Access:      c.Access,
		CreateBy:    c.CreateBy.Hex(),
		UpdatedAt:   c.UpdatedAt.Time(),
		Lessons:     ObjIDToHexID(c.Lessons),
	}
}
