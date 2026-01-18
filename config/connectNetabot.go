package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBNetabot *gorm.DB

func ConnectNetabot() {
	username := "root"
	password := "" // Kosongkan jika pakai XAMPP standar
	host := "127.0.0.1"
	port := "3306"
	dbName := "db_netabot"
	

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	fmt.Println("Koneksi Database Berhasil!", dbName)
	
	DBNetabot = database
}