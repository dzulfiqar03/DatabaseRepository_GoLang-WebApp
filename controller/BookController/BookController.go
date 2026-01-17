package BookController

import (
	"github.com/gin-gonic/gin"
	"database-repository/services/db-libraryis"
	"net/http"
)

func Index(c *gin.Context) {
	data, err,nama, dataList := dblibraryis.GetBook()
	

	if err != nil {
    // handle error
}
	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":"Book",
		"title": "Halaman Book",
		"data":  data,
		"nama":nama,
		"dataList": dataList,
	})
}
