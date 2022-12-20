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

var CourseDBInstance *courseDB

type courseDB struct {
	collection *mongo.Collection
}

func NewCourseDB(collectionName string) {
	CourseDBInstance = &courseDB{
		db.Collection(collectionName),
	}
}

func (c *courseDB) InsertCourse(course *core.Course) (err error) {
	ObjId, err := primitive.ObjectIDFromHex(course.CreateBy)
	if err != nil {
		return errs.InvalidUserID
	}
	courseDoc := courseDoc{
		ID:          primitive.NewObjectID(),
		Name:        course.Name,
		Description: course.Description,
		Access:      course.Access,
		CreateBy:    ObjId,
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		Lessons:     HexIDToObjID(course.Lessons),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = c.collection.InsertOne(ctx, courseDoc)
	if err != nil {
		return errs.InsertFailed
	}
	course.ID = courseDoc.ID.Hex()
	course.UpdatedAt = courseDoc.UpdatedAt.Time()
	return nil
}

func (c *courseDB) InsertManyCourse(courses []core.Course) (err error) {
	curriculumDocs := make([]interface{}, len(courses))
	for i, curriculum := range courses {
		ObjId, err := primitive.ObjectIDFromHex(curriculum.CreateBy)
		if err != nil {
			return errs.InvalidUserID
		}
		curriculumDocs[i] = courseDoc{
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
	for i := range courses {
		courses[i].ID = curriculumDocs[i].(courseDoc).ID.Hex()
		courses[i].UpdatedAt = curriculumDocs[i].(courseDoc).UpdatedAt.Time()
	}

	return nil
}

func (c *courseDB) FindCourse(id string) (core.Course, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return core.Course{}, errs.InvalidCourseID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var courseDoc courseDoc
	err = c.collection.FindOne(ctx, filter).Decode(&courseDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Course{}, errs.CourseNotFound
		}
		return core.Course{}, errs.FindFailed
	}
	return courseDoc.toCourse(), nil
}

func (c *courseDB) FindManyCourse(hexIDs []string) ([]core.Course, error) {
	objID := HexIDToObjID(hexIDs)
	filter := bson.M{"_id": bson.M{"$in": objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.FindFailed
	}
	var curriculumDocs []courseDoc
	if err = cursor.All(ctx, &curriculumDocs); err != nil {
		return nil, errs.FindFailed
	}
	courses := make([]core.Course, len(curriculumDocs))
	for i, curriculumDoc := range curriculumDocs {
		courses[i] = curriculumDoc.toCourse()
	}
	return courses, nil
}

func (c *courseDB) UpdateCourse(course *core.Course) error {
	objID, err := primitive.ObjectIDFromHex(course.ID)
	if err != nil {
		return errs.InvalidCourseID
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name":        course.Name,
			"description": course.Description,
			"access":      course.Access,
			"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
			"lessons":     HexIDToObjID(course.Lessons),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.UpdateTimeOut)
	defer cancel()
	result, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errs.UpdateFailed
	}
	if result.MatchedCount == 0 {
		return errs.CourseNotFound
	}
	return nil
}

func (c *courseDB) DeleteCourse(hexId string) error {
	objID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidCourseID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	if result.DeletedCount == 0 {
		return errs.CourseNotFound
	}
	return nil
}

func (c *courseDB) DeleteManyCourse(hexIds []string) error {
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

func (c *courseDB) Clear() {
	c.collection.DeleteMany(context.Background(), bson.M{})
}
