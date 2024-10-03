package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler()) // use promhttp
	http.ListenAndServe(":2112", nil)
}
