package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewCategory(name string, CreatedAt time.Time, UpdatedAt time.Time) *Category {
	return &Category{
		ID:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}
