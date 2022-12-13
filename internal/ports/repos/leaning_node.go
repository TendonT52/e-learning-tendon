package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type LearningNodeDB interface {
	InsertLeaningNodeDB(name, desc, acc, createBy string,
		node, next, prev []string) (core.LearningNode, error)
	GetLearningNodeById(id string) (core.LearningNode, error)
	DeleteLearningNode(hexId string) error
}
