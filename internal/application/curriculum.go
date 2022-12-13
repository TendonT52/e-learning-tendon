package application

import "github.com/TendonT52/e-learning-tendon/internal/ports/repos"

var CurriculumServiceInstance *curriculumService

type curriculumService struct {
	curriculumDB repos.CurriculumDB
}

func NewCurriculumService(curriculumDB repos.CurriculumDB) {
	CurriculumServiceInstance = &curriculumService{
		curriculumDB: curriculumDB,
	}
}

