package services

import (
	"log"
	"strconv"

	"github.com/VarunSharma3520/go-api/internal/models"
	"github.com/VarunSharma3520/go-api/internal/repository"
	"github.com/VarunSharma3520/go-api/internal/types"
	"github.com/VarunSharma3520/go-api/internal/utils"
)

// CreateTodoService converts DTO -> model and stores via repository.
func CreateTodoService(todo *types.Todo) {
	newTodo, err := utils.ConvertDTOtoModel(*todo)
	if err != nil {
		log.Println("CreateTodoService: convert error:", err)
		return
	}

	if err := repository.CreateTodo(newTodo); err != nil {
		log.Println("CreateTodoService: repository create error:", err)
		return
	}
}

// ReadTodoService reads from repository with skip & limit (strings -> ints)
func ReadTodoService(todo *types.Todo, skip string, limit string) []models.Todo {
	// Defaults
	skipInt := 0
	limitInt := 5

	// Convert query params if provided
	if skip != "" {
		if val, err := strconv.Atoi(skip); err == nil {
			skipInt = val
		} else {
			log.Println("ReadTodoService: invalid skip, using default 0:", err)
		}
	}
	if limit != "" {
		if val, err := strconv.Atoi(limit); err == nil {
			limitInt = val
		} else {
			log.Println("ReadTodoService: invalid limit, using default 5:", err)
		}
	}
	log.Println("ReadTodoService: using skip =", skipInt, "limit =", limitInt," from services")
	result, err := repository.ReadTodo(skipInt, limitInt)
	if err != nil {
		log.Println("ReadTodoService: repository read error:", err)
		return nil
	}
	return result
}

// UpdateTodoService patches via repository.
func UpdateTodoService(todo *types.Todo) []types.Todo {
	newTodo, err := utils.ConvertDTOtoModel(*todo)
	if err != nil {
		log.Println("UpdateTodoService: convert error:", err)
		return nil
	}

	if err := repository.PatchTodo(newTodo); err != nil {
		log.Println("UpdateTodoService: repository patch error:", err)
	}

	// No in-memory store anymore — return nil to preserve signature.
	return nil
}

// DeleteTodoService deletes by title in repository.
// Returns zero-value Todo (signature preserved).
func DeleteTodoService(title *string) types.Todo {
	if title == nil || *title == "" {
		log.Println("DeleteTodoService: empty title provided")
		return types.Todo{}
	}

	modelToDelete := models.Todo{
		Title: *title,
	}

	if err := repository.DeleteTodo(modelToDelete); err != nil {
		log.Println("DeleteTodoService: repository delete error:", err)
	}

	// No in-memory store to return the deleted item from — preserve signature.
	return types.Todo{}
}
