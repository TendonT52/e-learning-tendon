package core

type Lesson struct {
	ID          string
	Name        string
	Description string
	Access      string
	CreateBy    string
	UpdatedAt   string
	Nodes       []string
	NextLessons []string
	PrevLessons []string
}

// type BriefLesson struct {
// 	LearningNodeId          string
// 	LearningNodeName        string
// 	LearningNodeDescription string
// 	NextLearningNode        []string
// 	PrevLearningNode        []string
// }
