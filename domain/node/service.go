package node

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"gitlab.com/consensys-hellhound/pythia/event"
	"gitlab.com/consensys-hellhound/pythia/log"
	"math/big"
	"sort"
)

const (
	MinReputationScore = 1
	MaxReputationScore = 10
)

type ServiceMiddleware func(Service) Service

type Service interface {
	Save(*Node) error
	FindAll(context.Context) ([]Node, error)
	ElectForComputation(wantedNodes int, computationId uuid.UUID, code []byte) error
}

func NewService(nodes Repository, broker Broker) Service {
	return &service{
		nodes:  nodes,
		broker: broker,
	}
}

// build a Service interface chaining the middlewares
func Chain(s Service, mws ... ServiceMiddleware) Service {
	outer := s
	for _, mw := range mws {
		outer = mw(outer)
	}
	return outer
}

type service struct {
	nodes  Repository
	broker Broker
}

func (s service) Save(node *Node) error {
	if node.ReputationScore < MinReputationScore || node.ReputationScore > MaxReputationScore {
		return fmt.Errorf("reputation score must be between %d and %d", MinReputationScore, MaxReputationScore)
	}
	return s.nodes.Save(node)
}

func (s service) FindAll(_ context.Context) ([]Node, error) {
	return s.nodes.FindAll()
}

func (s service) ElectForComputation(wantedNodes int, computationId uuid.UUID, code []byte) (err error) {

	nodes, err := s.nodes.FindAll()
	if err != nil {
		return
	}
	log.Logger.Debugf("electing %d nodes", wantedNodes)
	electedNodes, err := elect(wantedNodes, nodes)

	log.Logger.Debugln("elected nodes")
	dump(electedNodes)

	var nodeIds []string

	for _, node := range electedNodes {
		nodeIds = append(nodeIds, node.NodeId)
	}

	cmd := event.HowlVmExecutionCommand{
		Type:          event.HowlVmExecutionCommandEvent,
		ComputationId: computationId.String(),
		Nodes:         nodeIds,
		Code:          hex.EncodeToString(code),
	}

	log.Logger.Debugln("publishing election result to broker")
	s.broker.Publish(event.TopicDomainEvent, cmd)

	return
}

func elect(number int, nodes []Node) (electedNodes []Node, err error) {
	if number > len(nodes)/2 {
		err = errors.New("cannot elect more than half of the network")
	}

	electedNodes = make([]Node, number)
	weightedNodes := make([]int, len(nodes))

	for index, node := range nodes {
		if index > 0 {
			weightedNodes[index] = weightedNodes[index-1] + node.ReputationScore
		} else {
			weightedNodes[index] = node.ReputationScore
		}
	}
	max := int64(weightedNodes[len(weightedNodes)-1])

	fmt.Println("\nMax index : ", max)
	fmt.Println("Weighted nodes : ", weightedNodes)

	for i := 0; i < number; {

		selector, _ := rand.Int(rand.Reader, big.NewInt(max))

		selectedNode := nodes[sort.Search(len(nodes)-1, func(i int) bool {
			return weightedNodes[i] > int(selector.Int64())
		})]
		fmt.Println("Selector : ", selector, "Selected node : ", selectedNode)

		if !contains(electedNodes, selectedNode) {
			electedNodes[i] = selectedNode
			i++
		}
	}

	return
}

func contains(nodes []Node, node Node) bool {
	for _, tmpNode := range nodes {
		if tmpNode.NodeId == node.NodeId {
			return true
		}
	}
	return false
}

func dump(in []Node) {
	for i, n := range in {
		p, _ := json.MarshalIndent(n, "", "\t")
		fmt.Printf("[ %d ] ==> %s", i, string(p))
	}
}
