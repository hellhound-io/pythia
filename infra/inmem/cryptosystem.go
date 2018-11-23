package inmem

import (
	"errors"
	"gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"
	"gitlab.com/consensys-hellhound/pythia/log"
	"sync"
)

type crypoStystemRepository struct {
	mtx           sync.RWMutex
	cryptosystems map[string]cryptosystem.CryptoSystem
}

func (r crypoStystemRepository) Store(cryptosystem cryptosystem.CryptoSystem) (cryptosystem.CryptoSystem, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.cryptosystems[cryptosystem.Id] = cryptosystem
	return cryptosystem, nil
}

func (r crypoStystemRepository) Find(id string) (*cryptosystem.CryptoSystem, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.cryptosystems[id]; ok {
		return &val, nil
	}
	return nil, errors.New("not found")
}

func (r crypoStystemRepository) FindByType(cryptosystemType cryptosystem.Type) ([]cryptosystem.CryptoSystem, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	cryptosystems := make([]cryptosystem.CryptoSystem, len(r.cryptosystems))

	for _, c := range r.cryptosystems {
		if c.Type == cryptosystemType {
			cryptosystems = append(cryptosystems, c)
		}
	}
	return cryptosystems, errors.New("not found")
}

func (r crypoStystemRepository) FindAll() ([]cryptosystem.CryptoSystem, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var cryptosystems []cryptosystem.CryptoSystem
	for _, c := range r.cryptosystems {
		cryptosystems = append(cryptosystems, c)
	}
	return cryptosystems, nil
}

func NewCryptoSystemRepository() cryptosystem.Repository {
	log.Logger.Infoln("starting in memory database")
	r := &crypoStystemRepository{
		cryptosystems: make(map[string]cryptosystem.CryptoSystem),
	}
	r.initData()
	return r
}

func (r *crypoStystemRepository) initData() {
	log.Logger.Debugln("init static data")
	r.cryptosystems = map[string]cryptosystem.CryptoSystem{
		"PHE_PAILLIER": {
			Id:      "PHE_PAILLIER",
			Name:    "PAILLIER",
			Type:    "Homomorphic",
			SubType: "PartiallyHomomorphic",
			Stage:   "Early",
		},
		"PHE_UNPADDED_RSA": {
			Id:      "PHE_UNPADDED_RSA",
			Name:    "UNPADDED RSA",
			Type:    "Homomorphic",
			SubType: "PartiallyHomomorphic",
			Stage:   "Early",
		},
		"PHE_ELGAMAL": {
			Id:      "PHE_ELGAMAL",
			Name:    "ELGAMAL",
			Type:    "Homomorphic",
			SubType: "PartiallyHomomorphic",
			Stage:   "Early",
		},
		"MPC_CRAMER_DAMGARD_NIELSER": {
			Id:      "MPC_CRAMER_DAMGARD_NIELSER",
			Name:    "CRAMER DAMGARD NIELSER",
			Type:    "MultiPartyComputation",
			SubType: "",
			Stage:   "Early",
		},
		"OBLIVIOUS_TRANSFER": {
			Id:      "OBLIVIOUS_TRANSFER",
			Name:    "OBLIVIOUS TRANSFER",
			Type:    "MultiPartyComputation",
			SubType: "",
			Stage:   "Early",
		},
		"SHAMIR_SECRET_SHARING": {
			Id:      "SHAMIR_SECRET_SHARING",
			Name:    "SHAMIR SECRET SHARING",
			Type:    "MultiPartyComputation",
			SubType: "",
			Stage:   "Early",
		},
	}
}

//	VALUES ('PHE_UNPADDED_RSA', 'UNPADDED RSA', 'Homomorphic', 'PartiallyHomomorphic','Early')
