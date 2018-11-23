package node

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/satori/go.uuid"
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
			"node",
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

func (s tracingService) Save(node *Node) (err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("node.save")
	err = s.Service.Save(node)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	}
	return

}

func (s tracingService) FindAll(ctx context.Context) (nodes []Node, err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("node.findall")
	nodes, err = s.Service.FindAll(ctx)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	}
	return
}

func (s tracingService) ElectForComputation(wantedNodes int, computationId uuid.UUID, code []byte) (err error) {
	var span opentracing.Span
	defer func() {
		span.Finish()
	}()
	span = opentracing.GlobalTracer().StartSpan("node.findall")
	span.SetTag("computationId", computationId.String())
	span.SetTag("wantedNodes", wantedNodes)
	err = s.Service.ElectForComputation(wantedNodes, computationId, code)
	if err != nil {
		span.SetTag("error", "true")
		span.SetTag("errorMessage", err.Error())
	}
	return
}
