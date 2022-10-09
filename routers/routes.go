package routers

import (
	"assignment-kedua/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	controller := controllers.DB{
		DB: db,
	}

	router.POST("/items", controller.CreateItem)
	router.GET("/items", controller.GetItems)
	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetOrders)
	router.PUT("/orders/:id", controller.UpdateOrder)
	router.DELETE("/orders/:id", controller.DeleteOrder)

	return router
}