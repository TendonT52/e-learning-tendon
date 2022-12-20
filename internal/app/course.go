package app

import (
	"errors"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
)

func CreateCourse(course *core.Course) error {
	_, err := reposInstance.UserDB.FindUser(course.CreateBy)
	if err != nil {
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"create by id is not invalid",
			)
		}
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"create by id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	lessons, err := reposInstance.LessonDB.FindManyLesson(course.Lessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find lessons in database",
		)
	}
	if len(lessons) != len(course.Lessons) {
		return errs.NewHttpError(
			http.StatusBadRequest,
			"some lessons are not found",
		)
	}
	err = reposInstance.CourseDB.InsertCourse(course)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't insert course to database",
		)
	}
	return nil
}

func GetCourse(id string) (core.Course, error) {
	course, err := reposInstance.CourseDB.FindCourse(id)
	if err != nil {
		if errors.Is(err, errs.InvalidCourseID) {
			return core.Course{}, errs.NewHttpError(
				http.StatusBadRequest,
				"course id is not invalid",
			)
		}
		if errors.Is(err, errs.CourseNotFound) {
			return core.Course{}, errs.NewHttpError(
				http.StatusBadRequest,
				"course id is not found",
			)
		}
		return core.Course{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find course in database",
		)
	}
	return course, nil
}

func UpdateCourse(course *core.Course) error {
	_, err := reposInstance.CourseDB.FindCourse(course.ID)
	if err != nil {
		if errors.Is(err, errs.InvalidCourseID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"course id is not invalid",
			)
		}
		if errors.Is(err, errs.CourseNotFound) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"course id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find course in database",
		)
	}
	lessons, err := reposInstance.LessonDB.FindManyLesson(course.Lessons)
	if err != nil {
		if errors.Is(err, errs.InvalidLessonID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"some lessons id are not invalid",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find lessons in database",
		)
	}
	if len(lessons) != len(course.Lessons) {
		return errs.NewHttpError(
			http.StatusBadRequest,
			"some lessons are not found",
		)
	}
	err = reposInstance.CourseDB.UpdateCourse(course)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't update course in database",
		)
	}
	return nil
}

func DeleteCourse(courseID string) error {
	_, err := reposInstance.CourseDB.FindCourse(courseID)
	if err != nil {
		if errors.Is(err, errs.InvalidCourseID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"course id is not invalid",
			)
		}
		if errors.Is(err, errs.CourseNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"course id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find course in database",
		)
	}
	err = reposInstance.CourseDB.DeleteCourse(courseID)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't delete course in database",
		)
	}
	return nil
}