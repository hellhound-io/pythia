package inmem

import (
	"encoding/json"
	"fmt"
	"gitlab.com/consensys-hellhound/pythia/domain/node"
	"testing"
)

func Test_nodeRepository_initData(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "nominal",
		},
	}



	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewNodeRepository()
			repostitory := r.(*nodeRepository)
			repostitory.initData()

			cs, _ := r.FindAll()
			for  _, c := range cs{
				fmt.Println("node : ", strNode(c))
			}
		})
	}
}

func strNode(c node.Node) string{
	p, _ := json.MarshalIndent(c, "", "\t")
	return string(p)
}
