package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gitlab.com/consensys-hellhound/cryptokit/paillier"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)

var (
	Key     *paillier.PrivateKey
	Encoder = paillier.NewPublicKeyEncoder()
)

func InitPaillier() {
	Key, _ = paillier.GenerateKey(rand.Reader, 128)
}

func PaillierKey(writer http.ResponseWriter, _ *http.Request) {
	bytes, err := Encoder.Encode(Key.PublicKey)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	writer.Write([]byte(hex.EncodeToString(bytes)))
}

func PaillierEncrypt(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	plain := string(body)
	plainInt, err := strconv.Atoi(plain)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	plainBig := big.NewInt(int64(plainInt))
	cipher, err := paillier.Encrypt(&Key.PublicKey, plainBig.Bytes())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	writer.Write([]byte(hex.EncodeToString(cipher)))
}

func PaillierDecrypt(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	cipher := string(body)
	cipherBytes, err := hex.DecodeString(cipher)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	decrypted, err := paillier.Decrypt(Key, cipherBytes)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	decryptedBig := new(big.Int)
	decryptedBig.SetBytes(decrypted)
	decryptedInt := int(decryptedBig.Int64())
	writer.Write([]byte(fmt.Sprintf("%d", decryptedInt)))
}
