package main

import (
	"github.com/alexgtn/go-middleware-metrics/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	metricsMiddleware := middleware.NewMetricsMiddleware()

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/venta/tv", tv).Methods(http.MethodGet)
	r.HandleFunc("/venta/computador", computador).Methods(http.MethodGet)

	r.Use(metricsMiddleware.Metrics)

	http.ListenAndServe(":8080", r)
}

func tv(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("TV"))
}

func computador(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Computador"))
}
