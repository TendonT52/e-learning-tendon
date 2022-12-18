package db

import (
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type nodeDoc struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `bson:"type"`
	Data      string             `bson:"data"`
	CreateBy  primitive.ObjectID `bson:"create_by"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

func (n *nodeDoc) toNode() core.Node {
	return core.Node{
		ID:        n.ID.Hex(),
		Type:      n.Type,
		Data:      n.Data,
		CreateBy:  n.CreateBy.Hex(),
		UpdatedAt: n.UpdatedAt.Time(),
	}
}
