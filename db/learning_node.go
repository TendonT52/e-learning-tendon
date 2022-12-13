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

var LearningNodeDBInstance *learningNodeDB

type learningNodeDB struct {
	collection *mongo.Collection
}

func NewLearningNodeDB(collectionName string) {
	LearningNodeDBInstance = &learningNodeDB{
		db.Collection(collectionName),
	}
}

func (ln *curriculumDB) InsertLeaningNodeDB(name, desc, acc, createBy string,
	node, next, prev []string) (core.LearningNode, error) {

	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return core.LearningNode{}, errs.ErrWrongFormat.From(err)

	}

	nodeObj, err := ArrayStringToArrayObjectId(node)
	if err != nil {
		return core.LearningNode{}, errs.ErrWrongFormat.From(err)
	}
	nextObj, err := ArrayStringToArrayObjectId(next)
	if err != nil {
		return core.LearningNode{}, errs.ErrWrongFormat.From(err)
	}

	prevObj, err := ArrayStringToArrayObjectId(next)
	if err != nil {
		return core.LearningNode{}, errs.ErrWrongFormat.From(err)
	}

	doc := LearningNodeDoc{
		ID:               primitive.NewObjectID(),
		Name:             name,
		Description:      desc,
		Access:           acc,
		CreateBy:         userID,
		UpdatedAt:        primitive.NewDateTimeFromTime(time.Now()),
		Node:             nodeObj,
		NextLearningNode: nextObj,
		PrevLearningNode: prevObj,
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = ln.collection.InsertOne(ctx, doc)
	if err != nil {
		return core.LearningNode{}, errs.ErrDatabase.From(err)
	}

	learningNode := core.LearningNode{
		ID:               doc.ID.Hex(),
		Name:             doc.Name,
		Description:      doc.Description,
		CreateBy:         doc.CreateBy.Hex(),
		Node:             node,
		NextLearningNode: next,
		PrevLearningNode: prev,
	}
	return learningNode, nil
}

func (ln *learningNodeDB) GetLearningNodeById(id string) (core.LearningNode, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.LearningNode{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := LearningNodeDoc{}
	err = ln.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.LearningNode{}, errs.ErrNotFound.From(err)
		}
		return core.LearningNode{}, errs.ErrDatabase.From(err)
	}

	learningNode := core.LearningNode{
		ID:               doc.ID.Hex(),
		Name:             doc.Name,
		Description:      doc.Description,
		CreateBy:         doc.CreateBy.Hex(),
		Node:             ArrayObjectIdToArrayString(doc.Node),
		NextLearningNode: ArrayObjectIdToArrayString(doc.NextLearningNode),
		PrevLearningNode: ArrayObjectIdToArrayString(doc.PrevLearningNode),
	}
	return learningNode, nil
}

func (ln *learningNodeDB) DeleteLearningNode(hexId string) error {
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.ErrInvalidToken.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := ln.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.ErrDatabase.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (ln *learningNodeDB) CleanUp() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	result, err := ln.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up leaning node collection, %v", err)
	}
	return int(result.DeletedCount)
}
