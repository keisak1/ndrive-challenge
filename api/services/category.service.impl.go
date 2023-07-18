package services

import (
	"api/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCategoryService(collection *mongo.Collection, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{collection, ctx}
}

func (ps *CategoryServiceImpl) CreateCategory(category *models.Category) (*models.Category, error) {

	category.CreatedAt = time.Now()
	category.UpdatedAt = category.CreatedAt

	res, err := ps.collection.InsertOne(ps.ctx, &category)

	if err != nil {
		return nil, err
	}

	var newCategory *models.Category
	query := bson.M{"_id": res.InsertedID}

	err = ps.collection.FindOne(ps.ctx, query).Decode(&newCategory)

	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (ps *CategoryServiceImpl) FindAll() ([]*models.Category, error) {
	collection := ps.collection.Database().Collection("Category")
	rows, err := collection.Find(context.Background(), bson.M{})
	var categories []*models.Category

	if err != nil {
		return nil, err
	} else {
		for rows.Next(context.Background()) {
			var category *models.Category
			if err := rows.Decode(&category); err != nil {
				return nil, err
			}
			categories = append(categories, category)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Process the products or return them as needed
	return categories, nil
}
