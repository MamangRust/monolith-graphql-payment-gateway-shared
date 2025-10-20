package observability

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-graphql-payment-gateway-pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TraceLoggerObservability provides tracing, logging, and request metrics.
type TraceLoggerObservability interface {
	// StartTracingAndLogging initializes tracing and logging for a given method.
	// It starts a span with optional attributes and logs the method start.
	// It returns the context with the span, the span, a function to end the span and record metrics,
	// the initial status of the operation, and a function to log success messages.
	//
	// Parameters:
	//   - ctx: The context for the traced method.
	//   - method: The name of the method to trace and log.
	//   - attrs: Optional attributes to add to the span.
	//
	// Returns:
	//   - context.Context: The context with the span.
	//   - trace.Span: The OpenTelemetry span for the traced method.
	//   - func(string): Function to end the span with a given status, recording metrics.
	//   - string: Initial status of the operation, defaulting to "success".
	//   - func(string, ...zap.Field): Function to log success messages with optional fields.
	StartTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (
		context.Context,
		trace.Span,
		func(string),
		string,
		func(string, ...zap.Field),
	)
	// RecordMetrics records a Prometheus metric for the given method and status.
	// It increments a counter and records the duration since the provided start time.
	RecordMetrics(method, status string, start time.Time)
}

// traceLoggerObservability provides tracing, logging, and request metrics.
type traceLoggerObservability struct {
	// Trace is the tracer used to create spans.
	Trace trace.Tracer

	// Logger is the structured logger for logging messages.
	Logger logger.LoggerInterface

	// RequestCounter counts incoming requests.
	RequestCounter *prometheus.CounterVec

	// RequestDuration records the duration of requests.
	RequestDuration *prometheus.HistogramVec
}

// NewTraceLoggerObservability creates a new traceLoggerObservability instance.
//
// It takes context.Context, trace.Tracer, logger.LoggerInterface, prometheus.CounterVec,
// and prometheus.HistogramVec as input.
//
// The returned traceLoggerObservability instance includes the input parameters and
// can be used to start tracing and logging, as well as recording metrics.
//
// If the input context is canceled, the tracing and logging will be stopped.
func NewTraceLoggerObservability(trace trace.Tracer, logger logger.LoggerInterface, counter *prometheus.CounterVec, duration *prometheus.HistogramVec) *traceLoggerObservability {
	return &traceLoggerObservability{
		Trace:           trace,
		Logger:          logger,
		RequestCounter:  counter,
		RequestDuration: duration,
	}
}

// StartTracingAndLogging initializes tracing and logging for a given method.
// It starts a span with optional attributes and logs the method start.
// It returns the context with the span, the span, a function to end the span and record metrics,
// the initial status of the operation, and a function to log success messages.
//
// Parameters:
//   - ctx: The context for the traced method.
//   - method: The name of the method to trace and log.
//   - attrs: Optional attributes to add to the span.
//
// Returns:
//   - context.Context: The context with the span.
//   - trace.Span: The OpenTelemetry span for the traced method.
//   - func(string): Function to end the span with a given status, recording metrics.
//   - string: Initial status of the operation, defaulting to "success".
//   - func(string, ...zap.Field): Function to log success messages with optional fields.
func (h *traceLoggerObservability) StartTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (
	context.Context,
	trace.Span,
	func(string),
	string,
	func(string, ...zap.Field),
) {
	start := time.Now()
	status := "success"

	ctx, span := h.Trace.Start(ctx, method)

	if len(attrs) > 0 {
		span.SetAttributes(attrs...)
	}

	span.AddEvent("Start: " + method)
	h.Logger.Info("Start: " + method)

	end := func(status string) {
		h.RecordMetrics(method, status, start)
		code := codes.Ok
		if status != "success" {
			code = codes.Error
		}
		span.SetStatus(code, status)
		span.End()
	}

	logSuccess := func(msg string, fields ...zap.Field) {
		span.AddEvent(msg)
		h.Logger.Info(msg, fields...)
	}

	return ctx, span, end, status, logSuccess
}

// RecordMetrics records a Prometheus metric for the given method and status.
// It increments a counter and records the duration since the provided start time.
func (h *traceLoggerObservability) RecordMetrics(method, status string, start time.Time) {
	if h.RequestCounter != nil {
		h.RequestCounter.WithLabelValues(method, status).Inc()
	}
	if h.RequestDuration != nil {
		h.RequestDuration.WithLabelValues(method, status).Observe(time.Since(start).Seconds())
	}
}
