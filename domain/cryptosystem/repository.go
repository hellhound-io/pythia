package cryptosystem

type Repository interface {
	Find(string) (*CryptoSystem, error)
	FindByType(Type) ([]CryptoSystem, error)
	FindAll() ([]CryptoSystem, error)
	Store(CryptoSystem) (CryptoSystem, error)
}
