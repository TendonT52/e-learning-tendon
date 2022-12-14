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

var LessonDBInstance *lessonDB

type lessonDB struct {
	collection *mongo.Collection
}

func NewLessonDB(collectionName string) {
	LessonDBInstance = &lessonDB{
		db.Collection(collectionName),
	}
}

func (l *lessonDB) InsertLessonDB(name, desc, acc, createBy string,
	node, next, prev []string) (core.Lesson, error) {

	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return core.Lesson{}, errs.ErrWrongFormat.From(err)
	}

	nodeObj, err := ArrayStringToArrayObjectId(node)
	if err != nil {
		return core.Lesson{}, errs.ErrWrongFormat.From(err)
	}
	nextObj, err := ArrayStringToArrayObjectId(next)
	if err != nil {
		return core.Lesson{}, errs.ErrWrongFormat.From(err)
	}

	prevObj, err := ArrayStringToArrayObjectId(next)
	if err != nil {
		return core.Lesson{}, errs.ErrWrongFormat.From(err)
	}

	doc := lessonDoc{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: desc,
		Access:      acc,
		CreateBy:    userID,
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		Node:        nodeObj,
		NextLesson:  nextObj,
		PrevLesson:  prevObj,
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = l.collection.InsertOne(ctx, doc)
	if err != nil {
		return core.Lesson{}, errs.ErrDatabase.From(err)
	}

	learningNode := core.Lesson{
		ID:          doc.ID.Hex(),
		Name:        doc.Name,
		Description: doc.Description,
		CreateBy:    doc.CreateBy.Hex(),
		Nodes:       node,
		NextLessons: next,
		PrevLessons: prev,
	}
	return learningNode, nil
}

func (l *lessonDB) InsertManyLessonDB(name, desc, acc []string,
	createBy string, node, next, prev [][]string) ([]core.Lesson, error) {

	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return []core.Lesson{}, errs.ErrWrongFormat.From(err)
	}

	var docs []interface{}
	for i := 0; i < len(name); i++ {
		nodeObj, err := ArrayStringToArrayObjectId(node[i])
		if err != nil {
			return []core.Lesson{}, errs.ErrWrongFormat.From(err)
		}
		nextObj, err := ArrayStringToArrayObjectId(next[i])
		if err != nil {
			return []core.Lesson{}, errs.ErrWrongFormat.From(err)
		}

		prevObj, err := ArrayStringToArrayObjectId(next[i])
		if err != nil {
			return []core.Lesson{}, errs.ErrWrongFormat.From(err)
		}

		doc := lessonDoc{
			ID:          primitive.NewObjectID(),
			Name:        name[i],
			Description: desc[i],
			Access:      acc[i],
			CreateBy:    userID,
			UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
			Node:        nodeObj,
			NextLesson:  nextObj,
			PrevLesson:  prevObj,
		}
		docs = append(docs, doc)
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = l.collection.InsertMany(ctx, docs)
	if err != nil {
		return []core.Lesson{}, errs.ErrDatabase.From(err)
	}

	var learningNodes []core.Lesson
	for _, doc := range docs {
		learningNode := core.Lesson{
			ID:          doc.(lessonDoc).ID.Hex(),
			Name:        doc.(lessonDoc).Name,
			Description: doc.(lessonDoc).Description,
			CreateBy:    doc.(lessonDoc).CreateBy.Hex(),
			Nodes:       ArrayObjectIdToArrayString(doc.(lessonDoc).Node),
			NextLessons: ArrayObjectIdToArrayString(doc.(lessonDoc).NextLesson),
			PrevLessons: ArrayObjectIdToArrayString(doc.(lessonDoc).PrevLesson),
		}
		learningNodes = append(learningNodes, learningNode)
	}
	return learningNodes, nil
}

func (l *lessonDB) GetLessonByID(hexID string) (core.Lesson, error) {
	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return core.Lesson{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := lessonDoc{}
	err = l.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Lesson{}, errs.ErrNotFound.From(err)
		}
		return core.Lesson{}, errs.ErrDatabase.From(err)
	}

	learningNode := core.Lesson{
		ID:          doc.ID.Hex(),
		Name:        doc.Name,
		Description: doc.Description,
		CreateBy:    doc.CreateBy.Hex(),
		Nodes:       ArrayObjectIdToArrayString(doc.Node),
		NextLessons: ArrayObjectIdToArrayString(doc.NextLesson),
		PrevLessons: ArrayObjectIdToArrayString(doc.PrevLesson),
	}
	return learningNode, nil
}

func (l *lessonDB) GetLessonManyByID(hexID []string) ([]core.Lesson, error) {
	objID, err := ArrayStringToArrayObjectId(hexID)
	if err != nil {
		return []core.Lesson{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objID}}}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	cursor, err := l.collection.Find(ctx, filter)
	if err != nil {
		return []core.Lesson{}, errs.ErrDatabase.From(err)
	}
	var docs []lessonDoc
	err = cursor.All(ctx, &docs)
	if err != nil {
		return []core.Lesson{}, errs.ErrDatabase.From(err)
	}
	if len(docs) == 0 {
		return []core.Lesson{}, errs.ErrNotFound.From(err)
	}
	var learningNodes []core.Lesson
	for _, doc := range docs {
		learningNode := core.Lesson{
			ID:          doc.ID.Hex(),
			Name:        doc.Name,
			Description: doc.Description,
			CreateBy:    doc.CreateBy.Hex(),
			Nodes:       ArrayObjectIdToArrayString(doc.Node),
			NextLessons: ArrayObjectIdToArrayString(doc.NextLesson),
			PrevLessons: ArrayObjectIdToArrayString(doc.PrevLesson),
		}
		learningNodes = append(learningNodes, learningNode)
	}
	return learningNodes, nil
}

func (l *lessonDB) DeleteLesson(hexID string) error {
	id, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := l.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.ErrDatabase.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (l *lessonDB) DeleteManyLesson(hexID []string) error {
	id, err := ArrayStringToArrayObjectId(hexID)
	if err != nil {
		return errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: id}}}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := l.collection.DeleteMany(ctx, filter)
	if err != nil {
		return errs.ErrDatabase.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (l *lessonDB) CleanUp() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	result, err := l.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up leaning node collection, %v", err)
	}
	log.Println()
	log.Printf("Learning node collection cleaned, %d records deleted", result.DeletedCount)
	return int(result.DeletedCount)
}
