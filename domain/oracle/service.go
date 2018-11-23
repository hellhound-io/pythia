package oracle

import "gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"

type Service interface {
	Advocate(Query) (QueryResponse, error)
}

func NewService(cryptoSystemService cryptosystem.Service) Service {
	return service{
		cryptoSystemService: cryptoSystemService,
	}
}

type service struct {
	cryptoSystemService cryptosystem.Service
}

// ServiceMiddleware is a middleware for the oracle.Service interface
type ServiceMiddleware func(Service) Service

// build a Service interface chaining the middlewares
func Chain(s Service, mws ... ServiceMiddleware) Service {
	outer := s
	for _, mw := range mws {
		outer = mw(outer)
	}
	return outer
}

func (s service) Advocate(query Query) (response QueryResponse, err error) {
	cryptoSystems, err := s.eligibleCryptoSystems(query)
	if err != nil {
		return
	}
	response, err = s.prepareQueryResponse(cryptoSystems)
	return
}

func (s service) eligibleCryptoSystems(query Query) (c []cryptosystem.CryptoSystem, err error) {
	var _type = cryptosystem.NoneType
	if !query.ComputationCriteria.Distributed{
		_type = cryptosystem.Homomorphic
	}
	cryptoSystemsByType, err := s.cryptoSystemService.FindByType(_type)
	if err != nil{
		return
	}
	c = cryptoSystemsByType
	return
}

func (s service) prepareQueryResponse(cryptoSystems []cryptosystem.CryptoSystem) (response QueryResponse, err error) {
	var recommendations []Recommendation
	for _, cryptoSystem := range cryptoSystems{
		recommendation := Recommendation{
			CryptoSystem: cryptoSystem,
		}
		recommendations = append(recommendations, recommendation)
	}
	response = QueryResponse{
		Recommendations: recommendations,
	}
	return
}
