package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name" binding:"required"`
	Price  float32            `json:"price" bson:"price" binding:"required"`
	Stock  int                `json:"stock" bson:"stock" binding:"required,min=1"`
	Rating int                `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
	Image  string             `json:"image" bson:"image"`
}

// Struct for Product Insertion to collection

type ProductInsert struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Price      float32            `json:"price" bson:"price" binding:"required"`
	Stock      int                `json:"stock" bson:"stock" binding:"required,min=1"`
	Rating     int                `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
	Image      string             `json:"image" bson:"image"`
	CategoryID []string           `json:"categoryId" bson:"categoryId" binding:required`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
type ProductEdit struct {
	ID     string  `json:"id" bson:"id"`
	Name   string  `json:"name" bson:"name" binding:"required"`
	Price  float32 `json:"price" bson:"price" binding:"required"`
	Stock  int     `json:"stock" bson:"stock" binding:"required,min=1"`
	Rating int     `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
	Image  string  `json:"image" bson:"image"`
}

// Struct for Product Deletion to collection
type ProductDelete struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

// DB Response to Product insert/search
type DBProductResponse struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name" binding:"required"`
	Price  float32            `json:"price" bson:"price" binding:"required"`
	Stock  int                `json:"stock" bson:"stock" binding:"required,min=1"`
	Rating int                `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
	Image  string             `json:"image" bson:"image"`
}

func NewProduct(name string, price float32, stock int, rating int, CategoryID []string, image string, CreatedAt time.Time, UpdatedAt time.Time) *ProductInsert {
	return &ProductInsert{
		ID:         primitive.NewObjectID(),
		Name:       name,
		Price:      price,
		Stock:      stock,
		Rating:     rating,
		CategoryID: CategoryID,
		Image:      image,
		CreatedAt:  CreatedAt,
		UpdatedAt:  UpdatedAt,
	}
}
