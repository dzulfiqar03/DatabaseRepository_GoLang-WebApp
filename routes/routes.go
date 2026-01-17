package routes

import (
	"database-repository/controller/BookController"
	"database-repository/controller/GiventController"
	"database-repository/controller/HomeController"
	"database-repository/controller/SeeuController"
	"database-repository/controller/SibanksaController"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", HomeController.Index)
	r.GET("/team", BookController.Index)
	r.GET("/db-sibanksa", SibanksaController.Index)
	r.GET("/db-seeU", SeeuController.Index)
	r.GET("/db-givent", GiventController.Index)
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Daftar Produk"})
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Login Berhasil"})
	})
}
