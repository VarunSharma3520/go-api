package main

import (
	"encoding/json"
	"github.com/VarunSharma3520/go-api/internal/config"
	"github.com/VarunSharma3520/go-api/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", routes.ApiV1Mux))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := config.Response{
			Status:  http.StatusText(200),
			Message: "Server healthy!",
			Data:    nil,
		}
		// 2️⃣ Set header to tell client it's JSON
		w.Header().Set("Content-Type", "application/json")

		// 3️⃣ Encode the Go struct/map into JSON and write to response
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	})

	log.Println("✅ Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
