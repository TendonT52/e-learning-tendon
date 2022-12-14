package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type LearningNodeDB interface {
	InsertLeaningNodeDB(name, desc, acc, createBy string,
		node, next, prev []string) (core.Lesson, error)
	GetLearningNodeById(id string) (core.Lesson, error)
	DeleteLearningNode(hexId string) error
}
