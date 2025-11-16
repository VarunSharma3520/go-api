package utils

import (
	"time"

	"github.com/VarunSharma3520/go-api/internal/models"
	"github.com/VarunSharma3520/go-api/internal/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDTOtoModel(dto types.Todo) (models.Todo, error) {

	t, err := time.Parse(time.RFC3339, dto.Reminder)
	if err != nil {
		return models.Todo{}, err
	}

	now := primitive.NewDateTimeFromTime(time.Now())

	return models.Todo{
		ID:          primitive.NewObjectID(),
		UserId:      dto.UserId,
		Title:       dto.Title,
		Description: dto.Description,
		Reminder:    primitive.NewDateTimeFromTime(t),
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
