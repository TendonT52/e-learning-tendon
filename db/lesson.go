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

var LessonDBInstance *lessonDB

type lessonDB struct {
	collection *mongo.Collection
}

func NewLessonDB(collectionName string) {
	LessonDBInstance = &lessonDB{
		db.Collection(collectionName),
	}
}

func (l *lessonDB) InsertLesson(lesson *core.Lesson) (err error) {
	userObjID, err := primitive.ObjectIDFromHex(lesson.CreateBy)
	if err != nil {
		return errs.InvalidUserID
	}
	lessonDoc := lessonDoc{
		ID:          primitive.NewObjectID(),
		Name:        lesson.Name,
		Description: lesson.Description,
		Access:      lesson.Access,
		CreateBy:    userObjID,
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		Nodes:       HexIDToObjID(lesson.Nodes),
		NextLessons: HexIDToObjID(lesson.NextLessons),
		PrevLessons: HexIDToObjID(lesson.PrevLessons),
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = l.collection.InsertOne(ctx, lessonDoc)
	if err != nil {
		return errs.InsertFailed
	}
	lesson.ID = lessonDoc.ID.Hex()
	lesson.UpdatedAt = lessonDoc.UpdatedAt.Time()
	return nil
}

func (l *lessonDB) InsertManyLesson(lessons []core.Lesson) (err error) {
	lessonDocs := make([]interface{}, len(lessons))
	for i, lesson := range lessons {
		userObjID, err := primitive.ObjectIDFromHex(lesson.CreateBy)
		if err != nil {
			return errs.InvalidUserID
		}
		lessonDocs[i] = lessonDoc{
			ID:          primitive.NewObjectID(),
			Name:        lesson.Name,
			Description: lesson.Description,
			Access:      lesson.Access,
			CreateBy:    userObjID,
			UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
			Nodes:       HexIDToObjID(lesson.Nodes),
			NextLessons: HexIDToObjID(lesson.NextLessons),
			PrevLessons: HexIDToObjID(lesson.PrevLessons),
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.InsertTimeOut)
	defer cancel()
	_, err = l.collection.InsertMany(ctx, lessonDocs)
	if err != nil {
		return errs.InsertFailed
	}
	for i := range lessons {
		lessons[i].ID = lessonDocs[i].(lessonDoc).ID.Hex()
		lessons[i].UpdatedAt = lessonDocs[i].(lessonDoc).UpdatedAt.Time()
	}
	return nil
}

func (l *lessonDB) FindLesson(hexID string) (core.Lesson, error) {
	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return core.Lesson{}, errs.InvalidLessonID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	var lessonDoc lessonDoc
	err = l.collection.FindOne(ctx, filter).Decode(&lessonDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.Lesson{}, errs.LessonNotFound
		}
		return core.Lesson{}, errs.FindFailed
	}
	return lessonDoc.toLesson(), nil
}

func (l *lessonDB) FindManyLesson(hexIDs []string) ([]core.Lesson, error) {
	objID := HexIDToObjID(hexIDs)
	filter := bson.M{"_id": bson.M{"$in": objID}}
	ctx, cancel := context.WithTimeout(context.Background(), config.FindTimeOut)
	defer cancel()
	cursor, err := l.collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.FindFailed
	}
	var lessonDocs []lessonDoc
	if err:= cursor.All(ctx, &lessonDocs); err != nil {
		return nil, errs.FindFailed
	}
	lessons := make([]core.Lesson, len(lessonDocs))
	for i, lessonDoc := range lessonDocs {
		lessons[i] = lessonDoc.toLesson()
	}
	return lessons, nil
}

func (l *lessonDB) UpdateLesson(lesson *core.Lesson) error {
	objID, err := primitive.ObjectIDFromHex(lesson.ID)
	if err != nil {
		return errs.InvalidLessonID
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"name":        lesson.Name,
		"description": lesson.Description,
		"access":      lesson.Access,
		"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
		"nodes":       HexIDToObjID(lesson.Nodes),
		"next_lessons": HexIDToObjID(lesson.NextLessons),
		"prev_lessons": HexIDToObjID(lesson.PrevLessons),
	}}
	ctx, cancel := context.WithTimeout(context.Background(), config.UpdateTimeOut)
	defer cancel()
	result, err := l.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errs.UpdateFailed
	}
	if result.MatchedCount == 0 {
		return errs.LessonNotFound
	}
	lesson.UpdatedAt = time.Now()
	return nil
}


func (l *lessonDB) DeleteLesson(hexId string) error {
	objID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return errs.InvalidLessonID
	}
	filter := bson.M{"_id": objID}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	result, err := l.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	if result.DeletedCount == 0 {
		return errs.LessonNotFound
	}
	return nil
}

func (l *lessonDB) DeleteManyLesson(hexIds []string) error {
	objIDs := HexIDToObjID(hexIds)
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), config.DeleteTimeOut)
	defer cancel()
	_, err := l.collection.DeleteMany(ctx, filter)
	if err != nil {
		return errs.DeleteFailed
	}
	return nil
}

func (l *lessonDB) Clear() {
	l.collection.DeleteMany(context.Background(), bson.M{})
}
