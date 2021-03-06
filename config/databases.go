package config

import (
	"backend_api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/belajar_backend?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Gagal koneksi database")
	}

	db.AutoMigrate(&models.Mahasiswa{}, &models.Dosen{})

	return db
}
