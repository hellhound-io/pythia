package oracle

import (
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

func (s tracingService) Advocate(query Query) (response QueryResponse, err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("pythia.oracle.advocate")
	span.SetTag("query", query.String())
	response, err = s.Service.Advocate(query)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	} else {
		span.SetTag("response", response.String())
	}
	return
}
