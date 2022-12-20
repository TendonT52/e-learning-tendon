package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type LessonDB interface {
	InsertLesson(lesson *core.Lesson) (err error)
	InsertManyLesson(lessons []core.Lesson) (err error)
	FindLesson(hexID string) (core.Lesson, error)
	FindManyLesson(hexIDs []string) ([]core.Lesson, error)
	UpdateLesson(lesson *core.Lesson) error
	DeleteLesson(hexId string) error
	DeleteManyLesson(hexIds []string) error
}
