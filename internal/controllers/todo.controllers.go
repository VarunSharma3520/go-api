package controllers

import (
	"encoding/json"
	// "log"
	"net/http"

	"github.com/VarunSharma3520/go-api/internal/config"
	"github.com/VarunSharma3520/go-api/internal/services"
	"github.com/VarunSharma3520/go-api/internal/types"
)

func CreateTodoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusMethodNotAllowed),
			Message: "Please check method request",
			Data:    nil,
		})
		return
	}

	var todo types.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid JSON body",
			Data:    nil,
		})
		return
	}

	// log.Printf("✅ Received todo Data: %+v", todo)

	services.CreateTodoService(&todo)
	json.NewEncoder(w).Encode(&config.Response{
		Status:  http.StatusText(http.StatusCreated),
		Message: "Todo created successfully",
	})
}

func ReadTodoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusMethodNotAllowed),
			Message: "Please check method request",
			Data:    nil,
		})
		return
	}
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")
	json.NewEncoder(w).Encode(config.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "Todos fetched successfully",
		Data:    services.ReadTodoService(&types.Todo{}, skip, limit),
	})
}

func UpdateTodoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusMethodNotAllowed),
			Message: "Please check method request",
			Data:    nil,
		})
		return
	}

	var todo types.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid JSON body",
			Data:    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(config.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "Todo updated successfully",
		Data:    services.UpdateTodoService(&todo),
	})
}

func DeleteTodoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&config.Response{
			Status:  http.StatusText(http.StatusMethodNotAllowed),
			Message: "Please check method request",
			Data:    nil,
		})
		return
	}

	title := r.URL.Query().Get("title")

	json.NewEncoder(w).Encode(config.Response{
		Status:  http.StatusText(http.StatusNoContent),
		Message: "Todo deleted successfully",
		Data:    services.DeleteTodoService(&title),
	})
}
