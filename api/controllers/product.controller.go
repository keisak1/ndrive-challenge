package controllers

import (
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
