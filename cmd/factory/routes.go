package factory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/letgomez/primecycle/internal/controller"
)

func InitRoutes() {
	orderController := controller.NewProductController()

	r := gin.Default()

	r.POST("/product", orderController.CreateProduct)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
