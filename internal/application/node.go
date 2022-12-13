package application

import "github.com/TendonT52/e-learning-tendon/internal/ports/repos"

var NodeServiceInstance *NodeService

type NodeService struct {
	NodeDB repos.NodeDB
}

func NewNodeService(NodeDB repos.NodeDB) {
	NodeServiceInstance = &NodeService{
		NodeDB: NodeDB,
	}
}
