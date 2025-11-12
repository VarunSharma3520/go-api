package services

import (
	"fmt"
	"strconv"

	"github.com/VarunSharma3520/go-api/internal/types"
)

var tododb = []types.Todo{}

func CreateTodoService(todo *types.Todo) {
	fmt.Printf("Creating todo: %+v\n", todo)
	tododb = append(tododb, *todo)

}

func ReadTodoService(todo *types.Todo, skip string, limit string) []types.Todo {
	fmt.Printf("Reading todo: %+v\n", todo)

	// Default values
	skipInt := 0
	limitInt := 5

	// Convert to int if present
	if skip != "" {
		if val, err := strconv.Atoi(skip); err == nil {
			skipInt = val
		}
	}
	if limit != "" {
		if val, err := strconv.Atoi(limit); err == nil {
			limitInt = val
		}
	}

	if skipInt > len(tododb) {
		return []types.Todo{}
	}
	if limitInt > len(tododb) {
		limitInt = len(tododb)
	}

	return tododb[skipInt:limitInt]
}

func UpdateTodoService(todo *types.Todo) []types.Todo {
	for i, v := range tododb {
		if v.Title == todo.Title {
			tododb[i] = *todo
			return tododb
		}
	}
	return nil
}

func DeleteTodoService(title *string) types.Todo {
	fmt.Printf("Deleting todo: %s\n", *title)
	for i, v := range tododb {
		if v.Title == *title {
			tododb = append(tododb[:i], tododb[i+1:]...)
			return v
		}
	}
	return types.Todo{}
}
