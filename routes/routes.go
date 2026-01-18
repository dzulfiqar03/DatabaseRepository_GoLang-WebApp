package routes

import (
	"database-repository/controller/GiventController"
	"database-repository/controller/HomeController"
	LibraryisController "database-repository/controller/LibraryIsController"
	"database-repository/controller/NetabotController"
	"database-repository/controller/SeeuController"
	"database-repository/controller/SibanksaController"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {



	
	r.GET("/", HomeController.Index)
	r.GET("/home", HomeController.Index)
	r.GET("/db-sibanksa", SibanksaController.Index)
	r.GET("/db-seeU", SeeuController.Index)
	r.GET("/db-givent", GiventController.Index)
	r.GET("/db-netabot", NetabotController.Index)
	r.GET("/db-libraryis", LibraryisController.Index)

	r.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Registrasi Berhasil"})
	})
	r.POST("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Logout Berhasil"})
	})
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Daftar Produk"})
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Login Berhasil"})
	})
}
