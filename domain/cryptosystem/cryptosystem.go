package cryptosystem

import "encoding/json"

// CryptoSystem entity
type CryptoSystem struct {
	Id		string `json:"id"`
	Name    string `json:"name"`
	Type    Type   `json:"type"`
	SubType string `json:"subType,omitempty"`
	Stage   Stage  `json:"stage,omitempty"`
}

// returns a new CryptoSystem
func New(id string, name string, setters ... Option) CryptoSystem {
	opts := DefaultOptions
	for _, setter := range setters {
		setter(opts)
	}
	return CryptoSystem{
		Id: id,
		Name:    name,
		Type:    opts.Type,
		SubType: opts.SubType,
		Stage:   opts.Stage,
	}
}

// Type describes the type of the cryptosystem
type Type string

// Valid cryptosystem types
const (
	Homomorphic           Type = "Homomorphic"
	MultiPartyComputation Type = "MultiPartyComputation"
	NoneType              Type = "None"
)

// Stage describes the stage of the cryptosystem
type Stage string

// Valid cryptosystem stages
const (
	Early     Stage = "Early"
	Advanced  Stage = "Advanced"
	SecondGen Stage = "SecondGen"
	NoneStage Stage = "None"
)

// Options define the options for the cryptosystem
type Options struct {
	Type    Type
	SubType string
	Stage   Stage
}

type Option func(*Options)

func WithStage(stage Stage) Option {
	return func(o *Options) {
		o.Stage = stage
	}
}

func WithType(cryptostemType Type) Option {
	return func(o *Options) {
		o.Type = cryptostemType
	}
}

func WithSubType(subType string) Option {
	return func(o *Options) {
		o.SubType = subType
	}
}

var DefaultOptions = &Options{
	Type:    NoneType,
	SubType: "",
	Stage:   NoneStage,
}

func (c CryptoSystem) String() string {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (c CryptoSystem) Bytes() []byte {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return nil
	}
	return b
}
