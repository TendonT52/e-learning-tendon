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

func (n *NodeDB) InsertNode(typ, data, createBy string) (core.Node, error) {
	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return core.Node{}, errs.ErrWrongFormat.From(err)
	}
	doc := nodeDoc{
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

func (n *NodeDB) InsertNodeMany(typ, data []string, createBy string) ([]core.Node, error) {
	if len(typ) != len(data) {
		return nil, errs.ErrWrongFormat
	}
	var docs []interface{}
	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return nil, errs.ErrWrongFormat.From(err)
	}
	for i := 0; i < len(typ); i++ {
		doc := nodeDoc{
			ID:        primitive.NewObjectID(),
			Type:      typ[i],
			Data:      data[i],
			CreateBy:  userID,
			UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		}
		docs = append(docs, doc)
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = n.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, errs.ErrDatabase.From(err)
	}
	var nodes []core.Node
	for _, doc := range docs {
		node := core.Node{
			ID:        doc.(nodeDoc).ID.Hex(),
			Type:      doc.(nodeDoc).Type,
			Data:      doc.(nodeDoc).Data,
			CreateBy:  doc.(nodeDoc).CreateBy.Hex(),
			UpdatedAt: doc.(nodeDoc).UpdatedAt.Time(),
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (n *NodeDB) GetNodeByID(id string) (core.Node, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Node{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := nodeDoc{}
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

func (n *NodeDB) GetNodeManyByID(id []string) ([]core.Node, error) {
	var objID []primitive.ObjectID
	for _, i := range id {
		obj, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return nil, errs.ErrWrongFormat.From(err)
		}
		objID = append(objID, obj)
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objID}}}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	cursor, err := n.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.ErrDatabase.From(err)
	}
	var docs []nodeDoc
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, errs.ErrDatabase.From(err)
	}
	if len(docs) == 0 {
		return nil, errs.ErrNotFound.From(err)
	}
	var nodes []core.Node
	for _, doc := range docs {
		node := core.Node{
			ID:        doc.ID.Hex(),
			Type:      doc.Type,
			Data:      doc.Data,
			CreateBy:  doc.CreateBy.Hex(),
			UpdatedAt: doc.UpdatedAt.Time(),
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (n *NodeDB) DeleteNodeByID(hexId string) error {
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

func (n *NodeDB) DeleteNodeManyByID(hexId []string) error {
	var objID []primitive.ObjectID
	for _, i := range hexId {
		obj, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return errs.ErrWrongFormat.From(err)
		}
		objID = append(objID, obj)
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objID}}}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := n.collection.DeleteMany(ctx, filter)
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
		log.Fatalf("Error while clean up node collection, %v", err)
	}
	log.Println("Clean up node collection", result.DeletedCount)
	return int(result.DeletedCount)
}
