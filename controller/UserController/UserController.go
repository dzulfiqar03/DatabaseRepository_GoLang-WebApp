package UserController

import (
	"github.com/gin-gonic/gin"
	"database-repository/services/UserService"
	"net/http"
)

func Index(c *gin.Context) {
	data, err,nama, dataList := userservice.GetUser()
	

	if err != nil {
    // handle error
}
	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":"Home",
		"title": "Halaman Home",
		"data":  data,
		"nama":nama,
		"dataList": dataList,
	})
}
