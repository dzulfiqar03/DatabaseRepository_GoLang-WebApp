package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    DBBook        *gorm.DB
    DBMember      *gorm.DB
    DBTransaction *gorm.DB
)

func ConnectLibraryIs() {
    username := "root"
    password := ""
    host     := "127.0.0.1"
    port     := "3306"

    // Helper untuk membuat koneksi
    createConn := func(dbName string) *gorm.DB {
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            username, password, host, port, dbName)
        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatalf("Gagal terhubung ke database %s: %v", dbName, err)
        }
        fmt.Println("Koneksi Berhasil:", dbName)
        return db
    }

    // Isi masing-masing variabel
    DBBook        = createConn("db_libraryis-book-service")
    DBMember      = createConn("db_libraryis-member-service")
    DBTransaction = createConn("db_libraryis-transaction-service")



	

}