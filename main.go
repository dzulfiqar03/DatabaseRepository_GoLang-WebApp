package main

import (
    "html/template"
	"github.com/gin-gonic/gin"
	"database-repository/config" // sesuaikan dengan nama module Anda
	"database-repository/routes"
)

func main() {
	// 1. Inisialisasi Database
	config.ConnectDatabase()

	r := gin.Default()

	r.Static("/assets", "./node_modules")

	r.Static("/styles", "frontend/resources/css/src")

	r.Static("/javascript", "frontend/resources/js")

    r.SetFuncMap(template.FuncMap{
    "add": func(a, b int) int {
        return a + b
    },
})

	r.LoadHTMLGlob("frontend/resources/views/**/*") 

	routes.Routes(r)

	r.Run(":8001")
}
