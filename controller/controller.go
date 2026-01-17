package controller

import (
	"database-repository/controller/BookController"
	"github.com/gin-gonic/gin"
)

func controller(c *gin.Context)  {
	BookController.Index(c)

}