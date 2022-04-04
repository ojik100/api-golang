package main

import (
	"backend_api/config"
	"backend_api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//models
	db := config.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "MOH.FAHRURROZI"})
	})

	//route mahasiswa
	r.GET("/mahasiswa", controllers.Readorder)
	r.POST("/mahasiswa", controllers.Addmahasiswa)
	r.PUT("/mahasiswa", controllers.Editmahasiswa)

	//route dosen
	r.GET("/dosen", controllers.Readdosen)

	r.Run()
}
