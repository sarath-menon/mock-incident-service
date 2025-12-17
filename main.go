package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "1.3.0"

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/logs", logsHandler)

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
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "# Basic metrics\nrequests_total 42\n")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"running","version":"%s","uptime":3600}`, version)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"logs":[{"level":"info","message":"Request received","timestamp":"2024-01-15T10:00:00Z"}]}`)
}
