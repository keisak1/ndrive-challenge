package services

import (
	"context"
	"time"

	"api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewProductService(collection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{collection, ctx}
}

func (ps *ProductServiceImpl) AddProduct(product *models.ProductInsert) (*models.DBProductResponse, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = product.CreatedAt
	product.Name = product.Name
	product.Stock = product.Stock
	product.Rating = product.Rating
	product.Image = product.Image

	res, err := ps.collection.InsertOne(ps.ctx, &product)

	if err != nil {
		return nil, err
	}

	var newProduct *models.DBProductResponse
	query := bson.M{"_id": res.InsertedID}

	err = ps.collection.FindOne(ps.ctx, query).Decode(&newProduct)

	if err != nil {
		return nil, err
	}

	return newProduct, nil

}

func (ps *ProductServiceImpl) DeleteProduct(product *models.ProductDelete) (*models.DBProductResponse, error) {

	//oid, _ := primitive.ObjectIDFromHex(product.ID)

	return nil, nil
}

func (ps *ProductServiceImpl) FindAll() ([]models.Product, error) {
	rows, err := ps.collection.Find(context.Background(), "Products")
	var products []models.Product

	if err != nil {
		return nil, err
	} else {

		for rows.Next(context.Background()) {
			var product models.Product
			if err := rows.Decode(&product); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Process the products or return them as needed
	return products, nil
}
