package httputil

import (
	"fmt"
	"gitlab.com/consensys-hellhound/pythia/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func StartServer(httpAddr string, errs chan error){
	log.Logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
	errs <- http.ListenAndServe(httpAddr, nil)
}


func HandleSigInt(errs chan error){
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
