package controllers

import (
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return CategoryController{categoryService}
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var category *models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var newCategory = models.NewCategory(category.Name, category.CreatedAt, category.UpdatedAt)
	result, err := cc.categoryService.CreateCategory(newCategory)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "created", "data": gin.H{"result": result}})

}
func (cc *CategoryController) FindAll(ctx *gin.Context) {
	categories, err := cc.categoryService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"status": "success", "data": gin.H{"categories": categories}})

}
