package repository

import (
	"context"
	"errors"
	"time"

	"github.com/VarunSharma3520/go-api/internal/db"
	"github.com/VarunSharma3520/go-api/internal/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	ErrNotFound = errors.New("todo not found")
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(dbName, collName string) *TodoRepository {
	// assumes db.Client is already initialized
	return &TodoRepository{
		collection: db.Client.Database(dbName).Collection(collName),
	}
}

// Insert a new todo
func (r *TodoRepository) Insert(ctx context.Context, todo *types.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, todo)
	return err
}

// Find returns todos matching non-empty fields of the filter.
// supports pagination via skip and limit (int values)
func (r *TodoRepository) Find(ctx context.Context, filter *types.Todo, skip, limit int) ([]types.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	bfilter := bson.M{}
	if filter != nil {
		if filter.UserId != "" {
			bfilter["userId"] = filter.UserId
		}
		if filter.Title != "" {
			bfilter["title"] = filter.Title
		}
		// add more filters as needed
	}

	findOptions := options.Find()
	if skip > 0 {
		findOptions.SetSkip(int64(skip))
	}
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	cur, err := r.collection.Find(ctx, bfilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []types.Todo
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Update by title (you can change to use ID or other unique key)
func (r *TodoRepository) UpdateByTitle(ctx context.Context, title string, todo *types.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"title": title}
	update := bson.M{"$set": bson.M{
		"userId":      todo.UserId,
		"title":       todo.Title,
		"description": todo.Description,
		"reminder":    todo.Reminder,
	}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

// Delete by title and return deleted document
func (r *TodoRepository) DeleteByTitle(ctx context.Context, title string) (*types.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"title": title}
	var deleted types.Todo
	err := r.collection.FindOneAndDelete(ctx, filter).Decode(&deleted)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
