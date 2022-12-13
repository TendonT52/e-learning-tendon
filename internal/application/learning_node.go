package application

import "github.com/TendonT52/e-learning-tendon/internal/ports/repos"

var LearningNodeServiceInstance *LeaningNodeService

type LeaningNodeService struct {
	leaningNodeDB repos.LearningNodeDB
}

func NewLeaningNodeService(learningNodeDB repos.LearningNodeDB) {
	LearningNodeServiceInstance = &LeaningNodeService{
		leaningNodeDB: learningNodeDB,
	}
}
