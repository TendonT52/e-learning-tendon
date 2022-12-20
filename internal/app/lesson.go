package app

import (
	"errors"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
)

func CreateLesson(lesson *core.Lesson) error {
	_, err := reposInstance.UserDB.FindUser(lesson.CreateBy)
	if err != nil {
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"create by id is not invalid",
			)
		}
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"create by id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	_, err = reposInstance.NodeDB.FindManyNode(lesson.Nodes)
	if err != nil {
		if errors.Is(err, errs.InvalidNodeID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some nodes id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find nodes in database",
		)
	}
	_, err = reposInstance.LessonDB.FindManyLesson(lesson.PrevLessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some prev lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find prev lessons in database",
		)
	}
	_, err = reposInstance.LessonDB.FindManyLesson(lesson.NextLessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some next lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find next lessons in database",
		)
	}
	err = reposInstance.LessonDB.InsertLesson(lesson)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't insert lesson to database",
		)
	}
	return nil
}

func GetLesson(lessonID string) (core.Lesson, error) {
	lesson, err := reposInstance.LessonDB.FindLesson(lessonID)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return core.Lesson{}, errs.NewHttpError(
				http.StatusBadRequest,
				"lesson id is not invalid",
			)
		}
		if errors.Is(err, errs.LessonNotFound) {
			return core.Lesson{}, errs.NewHttpError(
				http.StatusNotFound,
				"lesson id is not found",
			)
		}
		return core.Lesson{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find lesson in database",
		)
	}
	return lesson, nil
}

func UpdateLesson(lesson *core.Lesson) error {
	_, err := reposInstance.LessonDB.FindLesson(lesson.ID)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"lesson id is not invalid",
			)
		}
		if errors.Is(err, errs.LessonNotFound) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"lesson id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find lesson in database",
		)
	}
	_, err = reposInstance.NodeDB.FindManyNode(lesson.Nodes)
	if err != nil {
		if errors.Is(err, errs.InvalidNodeID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some nodes id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find nodes in database",
		)
	}
	_, err = reposInstance.LessonDB.FindManyLesson(lesson.PrevLessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some prev lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find prev lessons in database",
		)
	}
	_, err = reposInstance.LessonDB.FindManyLesson(lesson.NextLessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some next lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find next lessons in database",
		)
	}
	err = reposInstance.LessonDB.UpdateLesson(lesson)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't update lesson to database",
		)
	}
	return nil
}

func DeleteLesson(lessonID string) error {
	err := reposInstance.LessonDB.DeleteLesson(lessonID)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"lesson id is not invalid",
			)
		}
		if errors.Is(err, errs.LessonNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"lesson id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't delete lesson to database",
		)
	}
	return nil
}
