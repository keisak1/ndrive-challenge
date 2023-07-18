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
	router.POST("/create", pc.productController.AddProduct)
	router.DELETE("/delete/:id", pc.productController.Delete)
	router.GET("/:id", pc.productController.FindOne)
	router.PATCH("/edit/:id", pc.productController.EditProduct)
	router.GET("/", pc.productController.SearchProduct)
}
