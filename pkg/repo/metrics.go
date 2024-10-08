package repo

import "github.com/prometheus/client_golang/prometheus"

// Metrics contains general metrics for DAL methods.
type Metrics struct {
	callErrTotal *prometheus.CounterVec
	callDuration *prometheus.HistogramVec
}

// NewMetrics registers and returns common DAL metrics used by all
// services (namespace).
func NewMetrics(reg *prometheus.Registry, namespace, subsystem string, methodsFrom interface{}) (metric Metrics) {
	// metric.callErrTotal = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Namespace: namespace,
	// 		Subsystem: subsystem,
	// 		Name:      "errors_total",
	// 		Help:      "Amount of DAL errors.",
	// 	},
	// 	[]string{methodLabel},
	// )
	// reg.MustRegister(metric.callErrTotal)
	// metric.callDuration = prometheus.NewHistogramVec(
	// 	prometheus.HistogramOpts{
	// 		Namespace: namespace,
	// 		Subsystem: subsystem,
	// 		Name:      "call_duration_seconds",
	// 		Help:      "DAL call latency.",
	// 	},
	// 	[]string{methodLabel},
	// )
	// reg.MustRegister(metric.callDuration)

	// for _, methodName := range reflectx.MethodsOf(methodsFrom) {
	// 	l := prometheus.Labels{
	// 		methodLabel: methodName,
	// 	}
	// 	metric.callErrTotal.With(l)
	// 	metric.callDuration.With(l)
	// }

	return metric
}
