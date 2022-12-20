package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type NodeDB interface {
	InsertNode(node *core.Node) (err error)
	InsertManyNode(nodes []core.Node) (err error)
	FindNode(id string) (core.Node, error)
	FindManyNode(ids []string) ([]core.Node, error)
	UpdateNode(node *core.Node) error
	DeleteNode(hexId string) error
	DeleteManyNode(hexId []string) error
}
