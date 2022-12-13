package db

import (
	"context"
	"log"
	"time"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var NodeDBInstance *NodeDB

type NodeDB struct {
	collection *mongo.Collection
}

func NewNodeDB(collectionName string) {
	NodeDBInstance = &NodeDB{
		db.Collection(collectionName),
	}
}

func (n *NodeDB) InsertNodeDB(typ, data, createBy string) (core.Node, error) {
	userID, err := primitive.ObjectIDFromHex(createBy)
	doc := NodeDoc{
		ID:        primitive.NewObjectID(),
		Type:      typ,
		Data:      data,
		CreateBy:  userID,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = n.collection.InsertOne(ctx, doc)
	if err != nil {
		return core.Node{}, errs.ErrDatabase.From(err)
	}
	node := core.Node{
		ID:        doc.ID.Hex(),
		Type:      doc.Type,
		Data:      doc.Data,
		CreateBy:  doc.CreateBy.Hex(),
		UpdatedAt: doc.UpdatedAt.Time(),
	}
	return node, nil
}

func (n *NodeDB) GetNodeById(id string) (core.Node, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Node{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := NodeDoc{}
	err = n.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Node{}, errs.ErrNotFound.From(err)
		}
		return core.Node{}, errs.ErrDatabase.From(err)
	}
	node := core.Node{
		ID:        doc.ID.Hex(),
		Type:      doc.Type,
		Data:      doc.Data,
		CreateBy:  doc.CreateBy.Hex(),
		UpdatedAt: doc.UpdatedAt.Time(),
	}
	return node, nil
}

func (n *NodeDB) DeleteNode(hexId string) error {
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.ErrInvalidToken.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := n.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.ErrDatabase.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (n *NodeDB) CleanUp() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	result, err := n.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up leaning node collection, %v", err)
	}
	return int(result.DeletedCount)
}
