package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) CategoryAPI {
	return &categoryAPI{categoryService}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category data"})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to store category"})
		return
	}

	c.JSON(http.StatusOK, newCategory)
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	// Implement the UpdateCategory API endpoint
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	// Implement the DeleteCategory API endpoint
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	// Implement the GetCategoryByID API endpoint
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	// Implement the GetCategoryList API endpoint
}
