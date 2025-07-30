package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resp": "product created",
	})
}
