package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		// 固定采样、对所有数据都进行采样
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		// 是否启用 LoggingReporter、刷新缓冲区的频率、上报的 Agent 地址
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	// 根据配置项初始化 Tracer 对象，此处返回的是 opentracing.Tracer，并不是某个供应商的追踪系统的对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局的 Tracer 对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
