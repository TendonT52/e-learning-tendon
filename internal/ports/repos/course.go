package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type CurriculumDB interface {
	InsertCourse(curriculum *core.Course) (err error)
	InsertManyCourse(curriculums []core.Course) (err error)
	FindCourse(id string) (core.Course, error)
	FindManyCourse(hexIDs []string) ([]core.Course, error)
	UpdateCourse(curriculum *core.Course) error
	DeleteCourse(hexId string) error
	DeleteManyCourse(hexIds []string) error
}
