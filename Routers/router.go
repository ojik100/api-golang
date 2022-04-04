package Routes

import (
	"backend_api/controllers"

	"github.com/gin-gonic/gin"
)

func Stuprouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("mahasiswa", controllers.Readorder)
		v1.POST("mahasiswa", controllers.Addmahasiswa)
		v1.PUT("mahasiswa", controllers.Editmahasiswa)
	}

	return r
}
