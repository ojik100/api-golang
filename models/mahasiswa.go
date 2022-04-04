package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Mahasiswa struct {
	Id     string `gorm:"primaryKey"`
	Nim    string `json:"nim"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

func Getdata(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&db)
	c.JSON(http.StatusOK, gin.H{"data": db})
}
