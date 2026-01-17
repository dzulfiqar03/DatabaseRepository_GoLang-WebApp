package dblibraryis

import (
	"database-repository/config"
	"database-repository/model/Book"

)

func GetBook() ([]Book.Book, error, string, []string) {
    var books []Book.Book
    
    // Simpan hasil query ke variabel 'result'
    result := config.DBLibraryIs.Find(&books)
	nama:="Book Table"
	dataList:=[]string{"ali"}
    
    // Kembalikan datanya dan error-nya (jika ada)
    return books, result.Error, nama, dataList
}