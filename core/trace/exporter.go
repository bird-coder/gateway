/*
 * @Author: yujiajie
 * @Date: 2024-03-20 09:31:16
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:03:22
 * @FilePath: /Gateway/core/trace/exporter.go
 * @Description:
 */
package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

const DefaultService = "my-server"

// prometheus收集指标
func NewPromExporter(service string) error {
	exporter, err := prometheus.New()
	if err != nil {
		return err
	}
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(semconv.ServiceNameKey.String(service)),
	)
	if err != nil {
		return err
	}
	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(res),
	)
	otel.SetMeterProvider(provider)

	return nil
}

// jaeger收集指标
func NewJaegerExporter(service string, endpoint string) error {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return err
	}
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(service))
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(provider)
	return nil
}

// otlp收集指标
func NewOtlpExporter(service string, endpoint string) error {
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(service)))
	if err != nil {
		return err
	}
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(provider)
	return nil
}

// 控制台收集指标
func NewOsExporter(service string) error {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return err
	}
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(semconv.ServiceNameKey.String(service)),
	)
	if err != nil {
		return err
	}
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(provider)
	return nil
}
