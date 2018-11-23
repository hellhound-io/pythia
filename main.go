package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/satori/go.uuid"
	"gitlab.com/consensys-hellhound/pythia/domain/node"
	"gitlab.com/consensys-hellhound/pythia/domain/oracle"
	"gitlab.com/consensys-hellhound/pythia/event"
	"gitlab.com/consensys-hellhound/pythia/infra/inmem"
	"gitlab.com/consensys-hellhound/pythia/infra/nats"
	"os"

	"gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"
	"gitlab.com/consensys-hellhound/pythia/infra/http"
	"gitlab.com/consensys-hellhound/pythia/log"
	"net/http"
)

const (
	// listening address / port
	httpAddr = ":8080"
)

var (
	jaegerAgent = os.Getenv("JAEGER_AGENT_HOST_PORT")
	natsUrl     = os.Getenv("NATS_URL")
	nodeService node.Service
)

func main() {
	log.Logger.Infoln("starting Pythia")

	log.Logger.Debugln("launching components")

	log.Logger.Debugln("init nats broker")
	broker := nats.NewBroker(natsUrl)

	log.Logger.Debugln("intializing crypto system repository")
	// build a new cryptosystem.Repository
	//cryptoSystemRepository := cassandra.NewCryptoSystemRepository(os.Getenv("CASSANDRA_HOST"))
	cryptoSystemRepository := inmem.NewCryptoSystemRepository()

	// build a new cryptosystem.Service
	cryptoSystemService := cryptosystem.Chain(
		// base service
		cryptosystem.NewService(cryptoSystemRepository),
		// logging middleware
		cryptosystem.NewLoggingServiceMiddleware(),
		// tracing middleware
		cryptosystem.NewTracingServiceMiddleware(jaegerAgent),
		// instrumenting middleware
		cryptosystem.NewInstrumentingServiceMiddleware(),
	)

	// build a new oracle.Service
	oracleService := oracle.Chain(
		oracle.NewService(cryptoSystemService),
		oracle.NewLoggingServiceMiddleware(),
		oracle.NewTracingServiceMiddleware(jaegerAgent),
	)

	// build a new node.Repository
	nodeRepository := inmem.NewNodeRepository()
	// build a new node.Service
	nodeService = node.Chain(
		node.NewService(nodeRepository, broker),
		node.NewLoggingServiceMiddleware(),
		node.NewTracingServiceMiddleware(jaegerAgent),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", healthz)
	mux.Handle("/cryptosystem/", cryptosystem.MakeHandler(cryptoSystemService))
	mux.Handle("/oracle/", oracle.MakeHandler(oracleService))
	mux.Handle("/node/", node.MakeHandler(nodeService))

	/*
		DEMO purpose paillier functions
		TODO remove
	 */
	InitPaillier()
	mux.HandleFunc("/paillier/key", PaillierKey)
	mux.HandleFunc("/paillier/encrypt", PaillierEncrypt)
	mux.HandleFunc("/paillier/decrypt", PaillierDecrypt)
	/*
		END of demo paillier
	 */

	// Cross origin
	http.Handle("/", httputil.AccessControl(mux))
	// prometheus
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go httputil.StartServer(httpAddr, errs)
	go httputil.HandleSigInt(errs)

	log.Logger.Debugln("subscribing to domain events topic")
	broker.Subscribe(event.TopicDomainEvent, applyEvent)

	log.Logger.Log("terminated", <-errs)
}

func applyEvent(payload []byte) (err error) {
	defer func() {
		if err != nil {
			log.Logger.Errorf("applyEvent error : %s", err.Error())
		}
	}()
	log.Logger.Debugln("entering applyEvent with payload : ", string(payload))
	jsonParsed, err := gabs.ParseJSON(payload)
	commandType, ok := jsonParsed.Path("type").Data().(string)
	if !ok {
		err = errors.New("missing type field in event payload")
	}
	switch commandType {
	case event.VmExecutionRequestEvent:
		vmExecutionRequest := event.VmExecutionRequest{}
		err = json.Unmarshal(payload, &vmExecutionRequest)
		if err != nil {
			return
		}
		var computationId uuid.UUID
		var code []byte
		computationId, err = uuid.FromString(vmExecutionRequest.ComputationId)
		if err != nil {
			return
		}
		code, err = hex.DecodeString(vmExecutionRequest.Code)
		if err != nil {
			return
		}
		err = nodeService.ElectForComputation(vmExecutionRequest.WantedNodes, computationId, code)
	}
	return
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	data := map[string]string{
		"status": "UP",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
