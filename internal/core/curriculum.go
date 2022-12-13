package core

import "time"

const (
	PublicAccess    = "public"
	ProtectedAccess = "protected"
	PrivateAccess   = "private"
)

type Curriculum struct {
	ID           string
	Name         string
	Description  string
	Access       string
	CreateBy     string
	UpdatedAt    time.Time
	LearningNode []string
}

type BriefCurriculum struct {
	CurriculumId          string
	CurriculumName        string
	CurriculumDescription string
	Access                string
}
