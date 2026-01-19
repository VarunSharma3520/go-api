package repository

import (
	"context"
	"errors"
	// "log"
	"time"

	"github.com/VarunSharma3520/go-api/internal/db"
	"github.com/VarunSharma3520/go-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateTodo(todo models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	collection := db.Client.Database("todo").Collection("todos")
	_, err := collection.InsertOne(ctx, todo)
	return err
}

func ReadTodo(skip int, limit int) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	collection := db.Client.Database("todo").Collection("todos")
	filter := bson.M{}
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Todo
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func PatchTodo(todo models.Todo) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
    defer cancel()

    collection := db.Client.Database("todo").Collection("todos")

    // Filter by title
    filter := bson.M{"title": todo.Title}

    // Build update document dynamically (PATCH behavior)
    update := bson.M{}
    if todo.Description != "" {
        update["description"] = todo.Description
    }


    if len(update) == 0 {
        return errors.New("no fields to update")
    }

    _, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
    return err
}


func DeleteTodo(todo models.Todo) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
    defer cancel()
    collection := db.Client.Database("todo").Collection("todos")
    filter := bson.M{"title": todo.Title}
    _, err := collection.DeleteOne(ctx, filter)
    return err
}

