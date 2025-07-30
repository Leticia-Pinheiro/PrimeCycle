package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

type ProductRequest struct {
	Name string `json:"name"`
}

type ProductResponse struct {
	Message string `json:"message"`
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var reqBody ProductRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product created",
	})
}
