package oracle

import (
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

func (s loggingService) Advocate(query Query) (response QueryResponse, err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method: "oracleService.Advocate",
			"query":    query.String(),
			"response": response.String(),
			log.Took:   time.Since(begin),
			log.Err:    err,
		}).Info()
	}(time.Now())
	response, err = s.Service.Advocate(query)
	return
}
