package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})

}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	h := &Handler{store: NewStore()}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /incidents", h.listIncidents)
	mux.HandleFunc("POST /incidents", h.createIncident)

	logger.Info("servidor iniciado", "porta", 8080)

	if err := http.ListenAndServe(":8080", loggingMiddleware(logger, mux)); err != nil {
		logger.Error("falha ao iniciar o servidor", "error", err.Error())
		os.Exit(1)
	}
}
