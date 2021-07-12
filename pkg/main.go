package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imageservice/pkg/controllers"
	"github.com/imageservice/pkg/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/images", controllers.FindImages)
	r.GET("/images/:id", controllers.GetImage)
	r.POST("/images", controllers.CreateImage)
	r.PATCH("/images/:id", controllers.UpdateImage)

	r.Run(":8090")
}
