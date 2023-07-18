package controllers

import (
	"api/models"
	"api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return ProductController{productService}
}

func (pc *ProductController) AddProduct(ctx *gin.Context) {
	var product *models.ProductInsert

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result, err := pc.productService.AddProduct(product)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "created", "data": gin.H{"result": result}})

}

func (pc *ProductController) FindAll(ctx *gin.Context) {
	products, err := pc.productService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"status": "success", "data": gin.H{"products": products}})

}
func (pc *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	fmt.Println("This is the ID", id)
	result, err := pc.productService.DeleteProduct(id)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "delete", "data": gin.H{"result": result}})

}

func (pc *ProductController) FindOne(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	fmt.Println("This is the ID", id)
	result, err := pc.productService.FindOne(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"status": "success", "data": gin.H{"result": result}})

}

func (pc *ProductController) EditProduct(ctx *gin.Context) {
	var productEdit *models.ProductEdit
	id := ctx.Params.ByName("id")

	if err := ctx.ShouldBindJSON(&productEdit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	productEdit.ID = id

	result, err := pc.productService.EditProduct(productEdit)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"status": "success", "data": gin.H{"result": result}})
}

func (pc *ProductController) SearchProduct(ctx *gin.Context) {
	search := ctx.Query("search")

	result, err := pc.productService.SearchProduct(search)
	fmt.Println(result)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "No products found", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"status": "success", "data": gin.H{"result": result}})
}
