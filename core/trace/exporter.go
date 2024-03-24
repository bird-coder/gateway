/*
 * @Author: yujiajie
 * @Date: 2024-03-20 09:31:16
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-21 09:21:28
 * @FilePath: /gateway/core/trace/exporter.go
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
