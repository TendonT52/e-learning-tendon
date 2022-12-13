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

var CurriculumDBInstance *curriculumDB

type curriculumDB struct {
	collection *mongo.Collection
}

func NewCurriculumDB(collectionName string) {
	CurriculumDBInstance = &curriculumDB{
		db.Collection(collectionName),
	}
}

func (cu *curriculumDB) InsertCurriculumDB(name, desc, acc, createBy string, leaningNode []string) (core.Curriculum, error) {

	userID, err := primitive.ObjectIDFromHex(createBy)
	if err != nil {
		return core.Curriculum{}, errs.ErrWrongFormat.From(err)

	}

	leaningNodeObj, err := ArrayStringToArrayObjectId(leaningNode)

	doc := curriculumDoc{
		ID:           primitive.NewObjectID(),
		Name:         name,
		Description:  desc,
		Access:       acc,
		CreateBy:     userID,
		UpdatedAt:    primitive.NewDateTimeFromTime(time.Now()),
		LearningNode: leaningNodeObj,
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CreateTimeOut)
	defer cancel()
	_, err = cu.collection.InsertOne(ctx, doc)
	if err != nil {
		return core.Curriculum{}, errs.ErrDatabase.From(err)
	}

	cur := core.Curriculum{
		ID:           doc.ID.Hex(),
		Name:         doc.Name,
		Description:  doc.Description,
		Access:       doc.Access,
		CreateBy:     doc.ID.Hex(),
		UpdatedAt:    doc.UpdatedAt.Time(),
		LearningNode: leaningNode,
	}

	return cur, nil
}

func (cu *curriculumDB) GetCurriculumById(id string) (core.Curriculum, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Curriculum{}, errs.ErrWrongFormat.From(err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	doc := curriculumDoc{}
	err = cu.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Curriculum{}, errs.ErrNotFound.From(err)
		}
		return core.Curriculum{}, errs.ErrDatabase.From(err)
	}

	learningNodeHex := ArrayObjectIdToArrayString(doc.LearningNode)

	cur := core.Curriculum{
		ID:           doc.ID.Hex(),
		Name:         doc.Name,
		Description:  doc.Description,
		Access:       doc.Access,
		CreateBy:     doc.ID.Hex(),
		UpdatedAt:    doc.UpdatedAt.Time(),
		LearningNode: learningNodeHex,
	}
	return cur, nil
}

func (cu *curriculumDB) DeleteCurriculum(hexId string) error {
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.ErrInvalidToken.From(err)
	}
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeout)
	defer cancel()
	result, err := cu.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.ErrDatabase.From(err)
	}
	if result.DeletedCount == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (cu *curriculumDB) CleanUp() int {
	filter := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeout)
	defer cancel()
	result, err := cu.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Error while clean up curriculum collection, %v", err)
	}
	return int(result.DeletedCount)
}
