package cryptosystem

import "context"

// cryptosystem.Service interface
type Service interface {
	// finds a crypto system by id
	Find(string) (*CryptoSystem, error)
	// finds all crypto systems by type
	FindByType(Type) ([]CryptoSystem, error)
	// finds all crypto systems
	FindAll(context.Context) ([]CryptoSystem, error)
	// stores a crypto system
	Store(CryptoSystem) (CryptoSystem, error)
}

// ServiceMiddleware is a middleware for the cryptosystem.Service interface
type ServiceMiddleware func(Service) Service

// build a Service interface chaining the middlewares
func Chain(s Service, mws ... ServiceMiddleware) Service {
	outer := s
	for _, mw := range mws {
		outer = mw(outer)
	}
	return outer
}

// returns a new cryptosystem.Service
func NewService(repository Repository) Service {
	return &service{
		cryptosystems: repository,
	}
}

type service struct {
	// the repository to use for data access
	cryptosystems Repository
}

func (s service) Find(id string) (*CryptoSystem, error) {
	return s.cryptosystems.Find(id)
}

func (s service) FindByType(cryptosystemType Type) ([]CryptoSystem, error) {
	return s.cryptosystems.FindByType(cryptosystemType)
}

func (s service) FindAll(context.Context) ([]CryptoSystem, error) {
	return s.cryptosystems.FindAll()
}

func (s service) Store(c CryptoSystem) (CryptoSystem, error) {
	return s.cryptosystems.Store(c)
}
