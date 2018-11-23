package cryptosystem

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gitlab.com/consensys-hellhound/pythia/log"
	"io"
	"net/http"
)

func MakeHandler(s Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(log.Logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	storeHandler := kithttp.NewServer(
		makeStoreEndpoint(s),
		decodeStoreRequest,
		encodeResponse,
		opts...,
	)

	findHandler := kithttp.NewServer(
		makeFindEndpoint(s),
		decodeFindRequest,
		encodeResponse,
		opts...,
	)

	findAllHandler := kithttp.NewServer(
		makeFindAllEndpoint(s),
		decodeFindAllRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/cryptosystem/", storeHandler).Methods("POST")
	r.Handle("/cryptosystem/{id}/", findHandler).Methods("GET")
	r.Handle("/cryptosystem/", findAllHandler).Methods("GET")
	return r
}

func decodeStoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	cryptoSystem := CryptoSystem{}
	if err := json.NewDecoder(r.Body).Decode(&cryptoSystem); err != nil {
		return nil, err
	}
	return cryptoSystem, nil
}

func decodeFindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return vars["id"], nil
}

func decodeFindAllRequest(ctx context.Context, _ *http.Request) (interface{}, error) {
	return ctx, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	ApplicationJSON(w)
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errorType string
	var statusCode int
	switch errType := err.(type) {
	default:
		_ = errType
		errorType = "Default"
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	encodeJsonError(w, err, errorType)
}

func encodeJsonError(w io.Writer, err error, errType string) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
		"type":  errType,
	})
}

func ApplicationJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
