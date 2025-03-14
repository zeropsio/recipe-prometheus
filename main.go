package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheusRequestCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_counter",
		Help: "The total number of requests processed.",
	})
	var requestCounter uint64 = 0

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		prometheusRequestCounter.Inc()
		rc := atomic.AddUint64(&requestCounter, 1)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, strconv.FormatUint(rc, 10))
	})

	_ = http.ListenAndServe(":2112", nil)
}
