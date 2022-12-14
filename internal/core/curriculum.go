package core

import "time"


type Curriculum struct {
	ID          string
	Name        string
	Description string
	Access      string
	CreateBy    string
	UpdatedAt   time.Time
	Lessons     []string
}

// type BriefCurriculum struct {
// 	CurriculumId          string
// 	CurriculumName        string
// 	CurriculumDescription string
// 	Access                string
// }
