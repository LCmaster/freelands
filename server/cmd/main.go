package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status        string  `json:"status"`
	UptimeSeconds float64 `json:"uptimeSeconds"`
}

func main() {
	//TODO: Load .env
	//TODO: Parse command line arguments
	//TODO: Initialize logger
	//TODO: Initialize database
	fmt.Println("====> Starting HTTP server <====")
	startTime := time.Now()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		uptime := time.Since(startTime).Seconds()
		res := HealthResponse{
			Status:        "ok",
			UptimeSeconds: uptime,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(res)
	})
	fmt.Println("Server is running on http://localhost:4321")
	if err := http.ListenAndServe(":4321", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
