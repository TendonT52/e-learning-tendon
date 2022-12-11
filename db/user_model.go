package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Admin   = "admin"
	Teacher = "teacher"
	Student = "student"
)

const (
	PublicAccess    = "public"
	ProtectedAccess = "protected"
	PrivateAccess   = "private"
)

const (
	Required  = "required"
	Extension = "extension"
	Optional  = "optional"
)

const (
	Text  = "text"
	Image = "image"
	Video = "video"
	Pdf   = "pdf"
	Sound = "sound"
)

type userDoc struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Email        string             `bson:"email"`
	HashPassword string             `bson:"password"`
	Role         string             `bson:"role"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`
}

type Curriculum struct {
	Id          string    `bson:"_id,omitempty"`
	Name        string    `bson:"curriculum_name"`
	Description string    `bson:"description"`
	Access      string    `bson:"access"`
	CreateBy    string    `bson:"create_by"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type LearningNode struct {
	LearningNodeId          string
	LearningNodeName        string
	LearningNodeDescription string
	Node                    []Node
	NextLearningNode        []string
	PrevLearningNode        []string
}

type Node struct {
	Type      string
	NodeId    string
	Data      string
	Priority  string
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}
