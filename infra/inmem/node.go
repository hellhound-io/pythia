package inmem

import (
	"gitlab.com/consensys-hellhound/pythia/domain/node"
	"gitlab.com/consensys-hellhound/pythia/log"
	"sync"
)

type nodeRepository struct {
	mtx   sync.RWMutex
	nodes map[string]node.Node
}

func (r nodeRepository) Save(node *node.Node) (error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.nodes[node.NodeId] = *node
	return nil
}

func (r nodeRepository) FindAll() ([]node.Node, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var nodes []node.Node
	for _, n := range r.nodes {
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func NewNodeRepository() node.Repository {
	log.Logger.Infoln("starting node.Repository in memory database")
	r := &nodeRepository{
		nodes: make(map[string]node.Node),
	}
	r.initData()
	return r
}

func (r *nodeRepository) initData() {
	log.Logger.Debugln("init static data")
	r.nodes = map[string]node.Node{
		"00000000-0000-0000-0000-000000000001": {
			NodeId:          "00000000-0000-0000-0000-000000000001",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000002": {
			NodeId:          "00000000-0000-0000-0000-000000000002",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000003": {
			NodeId:          "00000000-0000-0000-0000-000000000003",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000004": {
			NodeId:          "00000000-0000-0000-0000-000000000004",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000005": {
			NodeId:          "00000000-0000-0000-0000-000000000005",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000006": {
			NodeId:          "00000000-0000-0000-0000-000000000006",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000007": {
			NodeId:          "00000000-0000-0000-0000-000000000007",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000008": {
			NodeId:          "00000000-0000-0000-0000-000000000008",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000009": {
			NodeId:          "00000000-0000-0000-0000-000000000009",
			Status:          "Active",
			ReputationScore: 5,
		},
		"00000000-0000-0000-0000-000000000010": {
			NodeId:          "00000000-0000-0000-0000-000000000010",
			Status:          "Active",
			ReputationScore: 5,
		},
	}
}
