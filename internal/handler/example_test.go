package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ankorstore/yokai-http-template/internal"
	"github.com/ankorstore/yokai/log/logtest"
	"github.com/ankorstore/yokai/trace/tracetest"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/fx"
)

func TestExampleHandler(t *testing.T) {
	var httpServer *echo.Echo
	var logBuffer logtest.TestLogBuffer
	var traceExporter tracetest.TestTraceExporter
	var metricsRegistry *prometheus.Registry

	internal.RunTest(t, fx.Populate(&httpServer, &logBuffer, &traceExporter, &metricsRegistry))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	httpServer.ServeHTTP(rec, req)

	// response assertion
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Welcome to http-app.")

	// logs assertion
	logtest.AssertHasLogRecord(
		t,
		logBuffer,
		map[string]interface{}{
			"level":   "info",
			"service": "http-app",
			"method":  "GET",
			"status":  200,
			"uri":     "/",
		},
	)

	// trace assertion
	tracetest.AssertHasTraceSpan(
		t,
		traceExporter,
		"GET /",
		semconv.HTTPMethod(http.MethodGet),
		semconv.HTTPRoute("/"),
		semconv.HTTPStatusCode(http.StatusOK),
	)

	// metrics assertion
	expectedMetric := `
		# HELP http_server_requests_total Number of processed HTTP requests
		# TYPE http_server_requests_total counter
		http_server_requests_total{method="GET",path="/",status="2xx"} 1
	`

	err := testutil.GatherAndCompare(
		metricsRegistry,
		strings.NewReader(expectedMetric),
		"http_server_requests_total",
	)
	assert.NoError(t, err)
}
