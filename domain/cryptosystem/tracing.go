package cryptosystem

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type tracingService struct {
	Service
}

func NewTracingServiceMiddleware(agentHostPort string) ServiceMiddleware {
	return func(s Service) Service {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: agentHostPort,
			},
		}
		tracer, _, err := cfg.New(
			"pythia",
		)
		if err != nil {
			panic(err)
		}
		opentracing.InitGlobalTracer(tracer)
		return tracingService{
			Service: s,
		}
	}
}

func (s tracingService) Find(id string) (cryptoSystem *CryptoSystem, err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("cryptosystem.find")
	span.SetTag("cryptoSystemId", id)
	cryptoSystem, err = s.Service.Find(id)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	} else {
		span.SetTag("cryptosystem", cryptoSystem.String())
	}
	return
}

func (s tracingService) FindAll(ctx context.Context) (cryptoSystems []CryptoSystem, err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("cryptosystem.findall")
	cryptoSystems, err = s.Service.FindAll(ctx)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	}
	return
}

func (s tracingService) Store(cryptoSystem CryptoSystem) (out CryptoSystem, err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("cryptosystem.store")
	out, err = s.Service.Store(cryptoSystem)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	} else {
		span.SetTag("cryptosystem", cryptoSystem.String())
	}
	return
}
