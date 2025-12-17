package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "1.1.0"

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/metrics", metricsHandler)

	log.Printf("Mock Incident Service v%s starting on :8080", version)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK - v%s", version)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mock Incident Service v%s", version)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "# Basic metrics\nrequests_total 42\n")
}
