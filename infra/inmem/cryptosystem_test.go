package inmem

import (
	"encoding/json"
	"fmt"
	"gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"
	"testing"
)

func Test_crypoStystemRepository_initData(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "nominal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewCryptoSystemRepository()
			repostitory := r.(*crypoStystemRepository)
			repostitory.initData()

			cs, _ := r.FindAll()
			for  _, c := range cs{
				fmt.Println("crytpo : ", str(c))
			}
		})
	}
}

func str(c cryptosystem.CryptoSystem) string{
	p, _ := json.MarshalIndent(c, "", "\t")
	return string(p)
}
