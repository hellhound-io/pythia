package cryptosystem

import (
	"context"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	namespace = "api"
	subsystem = "pythia"
)

var (
	fieldKeys = []string{"method"}
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentingServiceMiddleware() ServiceMiddleware {
	return func(s Service) Service {
		// counts the number of requests received
		requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys)
		// measure the duration of requests
		requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)
		return &instrumentingService{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			Service:        s,
		}
	}
}

func (s instrumentingService) Find(id string) (cryptoSystem *CryptoSystem, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "cryptosystemService.Find").Add(1)
		s.requestLatency.With("method", "cryptosystemService.Find").Observe(time.Since(begin).Seconds())
	}(time.Now())
	cryptoSystem, err = s.Service.Find(id)
	return
}

func (s instrumentingService) FindAll(ctx context.Context) (cryptoSystems []CryptoSystem, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "cryptosystemService.FindAll").Add(1)
		s.requestLatency.With("method", "cryptosystemService.FindAll").Observe(time.Since(begin).Seconds())
	}(time.Now())
	cryptoSystems, err = s.Service.FindAll(ctx)
	return
}

func (s instrumentingService) Store(cryptoSystem CryptoSystem) (out CryptoSystem, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "cryptosystemService.Store").Add(1)
		s.requestLatency.With("method", "cryptosystemService.Store").Observe(time.Since(begin).Seconds())
	}(time.Now())
	out, err = s.Service.Store(cryptoSystem)
	return
}
