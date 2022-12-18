package db

import (
	"context"
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

func (c *curriculumDB) InsertCurriculum(curriculum *core.Curriculum) (err error) {
	ObjId, err := primitive.ObjectIDFromHex(curriculum.CreateBy)
	if err != nil {
		return errs.InvalidUserID
	}
	curriculumDoc := curriculumDoc{
		ID:          primitive.NewObjectID(),
		Name:        curriculum.Name,
		Description: curriculum.Description,
		Access:      curriculum.Access,
		CreateBy:    ObjId,
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		Lessons:     HexIDToObjID(curriculum.Lessons),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = c.collection.InsertOne(ctx, curriculumDoc)
	if err != nil {
		return errs.InsertFailed
	}
	curriculum.ID = curriculumDoc.ID.Hex()
	curriculum.UpdatedAt = curriculumDoc.UpdatedAt.Time()
	return nil
}

func (c *curriculumDB) InsertManyCurriculum(curriculums []core.Curriculum) (err error) {
	curriculumDocs := make([]interface{}, len(curriculums))
	for i, curriculum := range curriculums {
		ObjId, err := primitive.ObjectIDFromHex(curriculum.CreateBy)
		if err != nil {
			return errs.InvalidUserID
		}
		curriculumDocs[i] = curriculumDoc{
			ID:          primitive.NewObjectID(),
			Name:        curriculum.Name,
			Description: curriculum.Description,
			Access:      curriculum.Access,
			CreateBy:    ObjId,
			UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
			Lessons:     HexIDToObjID(curriculum.Lessons),
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = c.collection.InsertMany(ctx, curriculumDocs)
	if err != nil {
		return errs.InsertFailed
	}
	for i := range curriculums {
		curriculums[i].ID = curriculumDocs[i].(curriculumDoc).ID.Hex()
		curriculums[i].UpdatedAt = curriculumDocs[i].(curriculumDoc).UpdatedAt.Time()
	}

	return nil
}

func (c *curriculumDB) FindCurriculum(id string) (core.Curriculum, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Curriculum{}, errs.InvalidCurriculumID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var curriculumDoc curriculumDoc
	err = c.collection.FindOne(ctx, filter).Decode(&curriculumDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Curriculum{}, errs.CurriculumNotFound
		}
		return core.Curriculum{}, errs.FindFailed
	}
	return curriculumDoc.toCurriculum(), nil
}

func (c *curriculumDB) FindManyCurriculum(hexIDs []string) ([]core.Curriculum, error) {
	objID := HexIDToObjID(hexIDs)
	filter := bson.M{"_id": bson.M{"$in": objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.FindFailed
	}
	var curriculumDocs []curriculumDoc
	if err = cursor.All(ctx, &curriculumDocs); err != nil {
		return nil, errs.FindFailed
	}
	curriculums := make([]core.Curriculum, len(curriculumDocs))
	for i, curriculumDoc := range curriculumDocs {
		curriculums[i] = curriculumDoc.toCurriculum()
	}
	return curriculums, nil

}

func (c *curriculumDB) UpdateCurriculum(curriculum *core.Curriculum) error {
	objID, err := primitive.ObjectIDFromHex(curriculum.ID)
	if err != nil {
		return errs.InvalidCurriculumID
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name":        curriculum.Name,
			"description": curriculum.Description,
			"access":      curriculum.Access,
			"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
			"lessons":     HexIDToObjID(curriculum.Lessons),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.UpdateTimeOut)
	defer cancel()
	result, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errs.UpdateFailed
	}
	if result.MatchedCount == 0 {
		return errs.CurriculumNotFound
	}
	return nil
}

func (c *curriculumDB) DeleteCurriculum(hexId string) error {
	objID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidCurriculumID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	if result.DeletedCount == 0 {
		return errs.CurriculumNotFound
	}
	return nil
}

func (c *curriculumDB) DeleteManyCurriculum(hexIds []string) error {
	objIDs := HexIDToObjID(hexIds)
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	_, err := c.collection.DeleteMany(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	return nil
}

func (c *curriculumDB) Clear(){
	c.collection.DeleteMany(context.Background(),bson.M{})
}