package core

import "time"

const (
	PublicAccess    = "public"
	ProtectedAccess = "protected"
	PrivateAccess   = "private"
)

type Curriculum struct {
	CurriculumId          string 
	CurriculumName        string 
	CurriculumDescription string
	BriefLearningNode     []BriefLearningNode
	Access                string             
	CreateBy              string             
	CreatedAt             time.Time          
	UpdatedAt             time.Time          
}

type BriefCurriculum struct {
	CurriculumId          string 
	CurriculumName        string 
	CurriculumDescription string
	Access                string 
}
