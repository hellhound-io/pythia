package node

type Repository interface {
	Save(*Node) error
	FindAll() ([]Node, error)
}
