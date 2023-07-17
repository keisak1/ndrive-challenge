package routes

import (
	"api/controllers"
	"api/services"

	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup, productService services.ProductService) {

	rg.GET("/products", pc.productController.FindAll)
	router := rg.Group("/products")
	router.DELETE("/delete/:id", pc.productController.Delete)

}
