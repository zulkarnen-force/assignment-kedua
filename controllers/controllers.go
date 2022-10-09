package controllers

import (
	"assignment-kedua/models"

	"fmt"
	"net/http"
	"strconv"

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

	
}


func (db *DB) UpdateOrder(ctx *gin.Context) {
	var orderUpdated models.Order
	ctx.ShouldBindJSON(&orderUpdated)

	id, _ := ctx.Params.Get("id")
	orderID, _ := strconv.Atoi(id)

	if err := db.DB.First(&orderUpdated, orderID).Error; err != nil {
	
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("order with id %s not found", id),
			"error_message": err.Error(),
		})
		return 
	
	}

	err := db.DB.Debug().Model(&orderUpdated).Where("order_id = ?", orderID).Updates(orderUpdated).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
		return 
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully updated data",
		"data": orderUpdated,
	})
}


func (db *DB) DeleteOrder(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	orderID, _ := strconv.Atoi(id)

	var order models.Order


	if err := db.DB.First(&order, orderID).Error; err != nil {
	
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("order with id %s not found", id),
			"error_message": err.Error(),
		})
		return 
	
	}


	err := db.DB.Debug().Delete(&order, orderID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":fmt.Sprintf("successfully deleted order with %d ", orderID),
	})
}