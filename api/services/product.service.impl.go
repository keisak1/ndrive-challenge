package services

import (
	"context"
	"fmt"
	"time"

	"api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ps *ProductServiceImpl) DeleteProduct(id string) (*mongo.DeleteResult, error) {
	collection := ps.collection.Database().Collection("Products")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectID}

	result, err := collection.DeleteOne(ps.ctx, query)

	return result, err
}

func (ps *ProductServiceImpl) FindAll() ([]*models.Product, error) {
	collection := ps.collection.Database().Collection("Products")

	// TODO:Sorting by price (ASCENDING)
	//filter := bson.D{}
	// opts := options.Find().SetSort(bson.D{{"price", 1}})

	// TODO:Sorting by price (DESCENDING)
	//filter := bson.D{}
	// opts := options.Find().SetSort(bson.D{{"price", -1}})
	rows, err := collection.Find(context.Background(), bson.M{})
	var products []*models.Product

	if err != nil {
		return nil, err
	} else {
		for rows.Next(context.Background()) {
			var product *models.Product
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

func (ps *ProductServiceImpl) FindOne(id string) (*models.Product, error) {
	collection := ps.collection.Database().Collection("Products")
	objectID, err := primitive.ObjectIDFromHex(id)
	var product *models.Product
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectID}
	collection.FindOne(ps.ctx, query).Decode(&product)
	return product, nil
}

func (ps *ProductServiceImpl) EditProduct(product *models.ProductEdit) (*models.ProductEdit, error) {
	collection := ps.collection.Database().Collection("Products")
	objectID, err := primitive.ObjectIDFromHex(string(product.ID))
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectID}
	var productEditing *models.Product
	collection.FindOne(ps.ctx, query).Decode(&productEditing)
	mergeProductFields(product, productEditing)
	update := bson.D{{Key: "$set", Value: bson.D{{"Name", product.Name}, {"Price", product.Price}, {"Stock", product.Stock}, {"Image", product.Image}}}}
	collection.UpdateOne(ps.ctx, query, update)
	return product, nil
}

func mergeProductFields(product *models.ProductEdit, productEditing *models.Product) {
	// Name
	if product.Name == "" {
		product.Name = productEditing.Name
	}

	// Price
	if product.Price == 0 {
		product.Price = productEditing.Price
	}

	// Stock
	if product.Stock == 0 {
		product.Stock = productEditing.Stock
	}

	// Image
	if product.Image == "" {
		product.Image = productEditing.Image
	}
	if product.Rating == 0 {
		product.Rating = productEditing.Rating
	}
}

func (ps *ProductServiceImpl) SearchProduct(search string) ([]*models.Product, error) {

	fmt.Println(search)
	collection := ps.collection.Database().Collection("Products")

	filter := bson.M{"name": bson.M{
		"$regex": primitive.Regex{
			Pattern: search,
			Options: "i",
		},
	}}
	cursor, err := collection.Find(ps.ctx, filter)

	if err != nil {
		return nil, err
	}
	var results []*models.Product

	for cursor.Next(ps.ctx) {
		var result *models.Product
		cursor.Decode(&result)
		fmt.Println(result)
		results = append(results, result)
	}

	return results, nil

}
