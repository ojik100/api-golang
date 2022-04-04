package controllers

import (
	"backend_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	Nim    string `json:"nim"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

//tampil data
func Readmahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

// order by id
func Readorder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Order("id desc, nim").Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": mhs})

}

//add data
func Addmahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	mhs := models.Mahasiswa{
		Nim:    dataInput.Nim,
		Nama:   dataInput.Nama,
		Alamat: dataInput.Alamat,
	}
	db.Create(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": mhs})

}

//funcition Edit
func Editmahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mhs models.Mahasiswa
	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ada"})
		return
	}

	//validasi
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&mhs).Update(dataInput)
	c.JSON(http.StatusOK, gin.H{"data": mhs})

	// delete data

}
