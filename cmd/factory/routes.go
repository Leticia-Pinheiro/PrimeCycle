package factory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/letgomez/primecycle/controller"
)

func InitRoutes() {
	r := gin.Default()

	r.POST("/product", controller.CreateProduct)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
