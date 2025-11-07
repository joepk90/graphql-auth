package stats

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	query = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "graphql_query_total",
		Help: "a counter of graphql queries",
	}, []string{"query_type", "method_name"})

	queryErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "graphql_query_error_total",
		Help: "a counter of graphql query errors",
	}, []string{"query_type", "method_name"})

	requestSize = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_size",
		Help: "request size",
	})

	responseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "response_time",
		Help: "response times",
	})

	responseSize = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_size",
		Help: "response size",
	})

	requestTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_total",
		Help: "total request",
	})
)

// PrometheusMetrics Provides global prometheus metrics
type PrometheusMetrics struct{}

// ObserveQuery observes a graphql query
func (p *PrometheusMetrics) ObserveQuery(queryType string, method string) {
	query.WithLabelValues(queryType, method).Inc()
}

// ObserveQueryError observes a graphql query error
func (p *PrometheusMetrics) ObserveQueryError(queryType string, method string) {
	queryErrors.WithLabelValues(queryType, method).Inc()
}

// ObserveRequestSize observes the request size
func (p *PrometheusMetrics) ObserveRequestSize(s float64) {
	requestSize.Add(s)
}

// ObserveResponseSize observes the response size
func (p *PrometheusMetrics) ObserveResponseSize(s float64) {
	responseSize.Add(s)
}

//ObserveResponseTime observers the time it took to complete the response
func (p *PrometheusMetrics) ObserveResponseTime(duration time.Duration) {
	responseTime.Observe(duration.Seconds())
}

// ObserveRequest observes a request
func (p *PrometheusMetrics) ObserveRequest() {
	requestTotal.Inc()
}
