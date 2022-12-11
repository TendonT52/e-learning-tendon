package core

type BriefLearningNode struct {
	LearningNodeId          string
	LearningNodeName        string
	LearningNodeDescription string
	NextLearningNode        []string
	PrevLearningNode        []string
}

type LearningNode struct {
	LearningNodeId          string
	LearningNodeName        string
	LearningNodeDescription string
	Node                    []Node
	Curriculum              []BriefCurriculum
	NextLearningNode        []string
	PrevLearningNode        []string
}
