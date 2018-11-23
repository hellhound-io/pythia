package oracle

import (
	"encoding/json"
	"gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"
)

type Query struct {
	CryptoSystemCriteria CryptoSystemCriteria `json:"cryptoSystem"`
	ComputationCriteria  ComputationCriteria  `json:"computation"`
}

type CryptoSystemCriteria struct {
	Encryption bool
}

type ComputationCriteria struct {
	Distributed    bool   `json:"distributed"`
	Interactive    bool   `json:"interactive"`
	TimeConstraint string `json:"timeConstraint"`
}

type QueryResponse struct {
	Recommendations []Recommendation `json:"recommendations"`
}

type Recommendation struct {
	CryptoSystem cryptosystem.CryptoSystem `json:"cryptoSystem"`
	Parameters   Parameters                `json:"parameters,omitempty"`
}

type Parameters struct {
	Values map[string]string `json:"values,omitempty"`
}

func (q Query) String() string {
	return string(q.Bytes())
}

func (q Query) Bytes() []byte {
	b, _ := json.MarshalIndent(q, "", "\t")
	return b
}

func (q QueryResponse) String() string {
	return string(q.Bytes())
}

func (q QueryResponse) Bytes() []byte {
	b, _ := json.MarshalIndent(q, "", "\t")
	return b
}
