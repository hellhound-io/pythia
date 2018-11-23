package cryptosystem

import (
	"context"
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

func (s loggingService) Find(id string) (cryptoSystem *CryptoSystem, err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method:       "cryptosystemService.Find",
			"cryptoSystemId": id,
			log.Took:         time.Since(begin),
			log.Err:          err,
		}).Info()
	}(time.Now())
	cryptoSystem, err = s.Service.Find(id)
	return
}

func (s loggingService) FindAll(ctx context.Context) (cryptoSystems []CryptoSystem, err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method: "cryptosystemService.FindAll",
			log.Took:   time.Since(begin),
			log.Err:    err,
		}).Info()
	}(time.Now())
	cryptoSystems, err = s.Service.FindAll(ctx)
	return
}

func (s loggingService) Store(cryptoSystem CryptoSystem) (out CryptoSystem, err error) {
	defer func(begin time.Time) {
		log.Logger.WithFields(logrus.Fields{
			log.Method:       "cryptosystemService.Store",
			"cryptoSystemId": cryptoSystem.Id,
			log.Took:         time.Since(begin),
			log.Err:          err,
		}).Info()
	}(time.Now())
	out, err = s.Service.Store(cryptoSystem)
	return
}
