package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

// Inicializa o Tracer do OpenTelemetry
func initTracer() (*sdktrace.TracerProvider, error) {
	// Cria um exportador OTLP (envia traces para um coletor)
	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint("localhost"), // Coletor local (ex: Jaeger/OTel Collector)
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar exportador OTLP: %v", err)
	}

	// Configura o provedor de traces
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("lambda-go"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	tracer = tp.Tracer("github.com/sibelly/aws-lambda-example")
	return tp, nil
}

// Handler da Lambda
func handler(ctx context.Context) (string, error) {
	// Cria um span customizado
	ctx, span := tracer.Start(ctx, "handler-execution")
	defer span.End()

	log.Println("Lambda executada com sucesso!")
	return "Olá, Mundo! Traces enviados para OpenTelemetry.", nil
}

func handler2(ctx context.Context) (string, error) {

	log.Println("Lambda executada com sucesso!")
	return "Olá, Mundo! Traces enviados para OpenTelemetry.", nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Erro ao inicializar OpenTelemetry: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Erro ao desligar tracer: %v", err)
		}
	}()

	lambda.Start(handler)
}
