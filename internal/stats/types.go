package stats

import "time"

// Metrics provides an interface for stats
type Metrics interface {
	ObserveQuery(queryType string, method string)
	ObserveQueryError(queryType string, method string)
	ObserveRequestSize(float64)
	ObserveResponseSize(float64)
	ObserveResponseTime(time.Duration)
	ObserveRequest()
}
