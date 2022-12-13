package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type CurriculumDB interface {
	InsertCurriculumDB(name, desc, acc, createBy string, leaningNode []string) (core.Curriculum, error)
	GetCurriculumById(id string) (core.Curriculum, error)
	DeleteCurriculum(hexId string) error
}
