package models

import (
	"context"
	"fmt"
	"time"

	"github.com/VarunSharma3520/go-api/internal/db"
	"github.com/VarunSharma3520/go-api/internal/types"
)

func InsertTodo(todo types.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("todo").Collection("todos")
	_, err := collection.InsertOne(ctx, todo)
	fmt.Println("Inserted Todo",collection)
	return err
}
