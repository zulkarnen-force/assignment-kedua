package controllers

import (
	"assignment-kedua/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
} 

func (db *DB) CreateItem(ctx *gin.Context) {
	var newItem models.Item

	ctx.ShouldBindJSON(&newItem)

	err := db.DB.Debug().Create(&newItem).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
		return 
	}

	ctx.JSON(http.StatusCreated, newItem)
}


func (db *DB) CreateOrder(ctx *gin.Context) {
	var order models.Order
	ctx.ShouldBindJSON(&order)

	err := db.DB.Debug().Create(&order).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
		return 
	}
	
	ctx.JSON(http.StatusCreated, order)
}

func (db *DB) GetItems(ctx *gin.Context) {

	var items []models.Item
	db.DB.Find(&items)
	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}


func (db *DB) GetOrders(ctx *gin.Context) {

	var orders []models.Order

	err := db.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error


	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message":err,
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})

	return
}