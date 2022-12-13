package core

type LearningNode struct {
	ID               string
	Name             string
	Description      string
	CreateBy         string
	Node             []string
	NextLearningNode []string
	PrevLearningNode []string
}

type BriefLearningNode struct {
	LearningNodeId          string
	LearningNodeName        string
	LearningNodeDescription string
	NextLearningNode        []string
	PrevLearningNode        []string
}
