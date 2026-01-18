package dblibraryis

import (
	"database-repository/config"
	"database-repository/model/dbLibraryis"
)

func GetBooks() ([]dblibraryis.Book, error, string) {
	var book []dblibraryis.Book

	nama := "Books Table"

	result := config.DBBook.
		Table("books").
		Preload("Book_Detail").
		Find(&book)

	return book, result.Error, nama
}