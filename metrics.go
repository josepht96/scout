package main

import "github.com/prometheus/client_golang/prometheus"

var labels = []string{"src_node", "src_pod", "dest_node", "dest_pod"}

func createMetricTotalRequests() *prometheus.CounterVec {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "scout_total_requests",
			Help: "counts the number of requests",
		}, labels,
	)
	prometheus.MustRegister(counter)

	return counter
}

func createMetricTotalLatency() *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "scout_total_latency",
			Help: "gauge for total latency ms",
		}, labels,
	)
	prometheus.MustRegister(gauge)

	return gauge
}
func createMetricDNSDur() *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "scout_total_dns_duration",
			Help: "gauge for total dns duration",
		}, labels,
	)
	prometheus.MustRegister(gauge)

	return gauge
}

func createMetricConnDur() *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "scout_total_conn_duration",
			Help: "gauge for total connection creation duration",
		}, labels,
	)
	prometheus.MustRegister(gauge)

	return gauge
}

func createMetricServerProcessingDur() *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "scout_total_server_processing_duration",
			Help: "gauge for total server processing duration",
		}, labels,
	)
	prometheus.MustRegister(gauge)

	return gauge
}
