package routes

import (
	"api/controllers"
	"api/services"

	"github.com/gin-gonic/gin"
)

type CategoryRouteController struct {
	categoryController controllers.CategoryController
}

func NewRouteCategoryController(categoryController controllers.CategoryController) CategoryRouteController {
	return CategoryRouteController{categoryController}
}

func (cc *CategoryRouteController) CategoryRoute(rg *gin.RouterGroup, categoryService services.CategoryService) {

	rg.GET("/categories", cc.categoryController.FindAll)
	rg.POST("/categories", cc.categoryController.CreateCategory)
}
