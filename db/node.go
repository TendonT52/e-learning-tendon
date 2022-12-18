package db

import (
	"context"
	"errors"
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

func (n *NodeDB) InsertNode(node *core.Node) (err error) {
	objID, _ := primitive.ObjectIDFromHex(node.CreateBy)
	nodeDoc := nodeDoc{
		ID:        primitive.NewObjectID(),
		Type:      node.Type,
		Data:      node.Data,
		CreateBy:  objID,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = n.collection.InsertOne(ctx, nodeDoc)
	if err != nil {
		return err
	}
	node.ID = nodeDoc.ID.Hex()
	node.UpdatedAt = nodeDoc.UpdatedAt.Time()
	return nil
}

func (n *NodeDB) InsertManyNode(nodes []core.Node) (err error) {
	nodeDocs := make([]interface{}, len(nodes))
	for i, node := range nodes {
		objID, _ := primitive.ObjectIDFromHex(node.CreateBy)
		nodeDocs[i] = nodeDoc{
			ID:        primitive.NewObjectID(),
			Type:      node.Type,
			Data:      node.Data,
			CreateBy:  objID,
			UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = n.collection.InsertMany(ctx, nodeDocs)
	if err != nil {
		return err
	}
	for i := range nodes {
		nodes[i].ID = nodeDocs[i].(nodeDoc).ID.Hex()
		nodes[i].UpdatedAt = nodeDocs[i].(nodeDoc).UpdatedAt.Time()
	}
	return nil
}

func (n *NodeDB) FindNode(id string) (core.Node, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Node{}, errs.InvalidNodeID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var nodeDoc nodeDoc
	err = n.collection.FindOne(ctx, filter).Decode(&nodeDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Node{}, errs.NodeNotFound
		}
		return core.Node{}, errs.FindFailed
	}
	node := nodeDoc.toNode()
	return node, nil
}

func (n *NodeDB) FindManyNode(ids []string) ([]core.Node, error) {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, errs.InvalidNodeID.From(errors.New("invalid id: " + id))
		}
		objIDs[i] = objID
	}
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	cursor, err := n.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.FindFailed
	}
	var nodeDocs []nodeDoc
	if err = cursor.All(ctx, &nodeDocs); err != nil {
		return nil, errs.FindFailed
	}
	nodes := make([]core.Node, len(nodeDocs))
	for i, nodeDoc := range nodeDocs {
		nodes[i] = nodeDoc.toNode()
	}
	return nodes, nil
}

func (n *NodeDB) UpdateNode(node *core.Node) error {
	objID, err := primitive.ObjectIDFromHex(node.ID)
	if err != nil {
		return errs.InvalidNodeID
	}
	node.UpdatedAt = time.Now()
	update := bson.M{"$set": bson.M{
		"type":       node.Type,
		"data":       node.Data,
		"updated_at": primitive.NewDateTimeFromTime(node.UpdatedAt),
	}}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.UpdateTimeOut)
	defer cancel()
	result, err := n.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errs.UpdateFailed
	}
	if result.MatchedCount == 0 {
		return errs.NodeNotFound
	}
	return nil
}

func (n *NodeDB) DeleteNode(hexId string) error {
	objID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidNodeID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := n.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	if result.DeletedCount == 0 {
		return errs.NodeNotFound
	}
	return nil
}

func (n *NodeDB) DeleteManyNode(hexId []string) error {
	objIDs := HexIDToObjID(hexId)
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	_, err := n.collection.DeleteMany(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	return nil
}

func (n *NodeDB) Clear() {
	n.collection.DeleteMany(context.Background(), bson.M{})
}
