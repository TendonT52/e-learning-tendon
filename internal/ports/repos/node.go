package repos

import "github.com/TendonT52/e-learning-tendon/internal/core"

type NodeDB interface {
	InsertNodeDB(typ, data, createBy string) (core.Node, error)
	GetNodeById(id string) (core.Node, error)
	DeleteNode(hexId string) error
}
