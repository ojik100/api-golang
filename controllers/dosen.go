package controllers

import (
	"backend_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Readdosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen []models.Dosen

	db.Order("id desc, nim").Find(&dosen)
	c.JSON(http.StatusOK, gin.H{"data": dosen})

}
