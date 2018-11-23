package node

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/consensys-hellhound/pythia/log"
	"time"
)

type loggingService struct {
	Service
}

func NewLoggingServiceMiddleware() ServiceMiddleware {
	return func(s Service) Service {
		return loggingService{
			Service: s,
		}
	}
}

func (s loggingService) Save(node *Node) (err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method: "nodeService.Save",
			"nodeId":   node.NodeId,
			log.Took:   time.Since(begin),
			log.Err:    err,
		}).Info()
	}(time.Now())
	err = s.Service.Save(node)
	return

}

func (s loggingService) FindAll(ctx context.Context) (nodes []Node, err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method: "nodeService.FindAll",
			log.Took:   time.Since(begin),
			log.Err:    err,
		}).Info()
	}(time.Now())
	nodes, err = s.Service.FindAll(ctx)
	return
}

func (s loggingService) ElectForComputation(wantedNodes int, computationId uuid.UUID, code []byte) (err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method:      "nodeService.ElectForComputation",
			"computationId": computationId.String(),
			"wantedNodes":   wantedNodes,
			log.Took:        time.Since(begin),
			log.Err:         err,
		}).Info()
	}(time.Now())
	err = s.Service.ElectForComputation(wantedNodes, computationId, code)
	return
}
