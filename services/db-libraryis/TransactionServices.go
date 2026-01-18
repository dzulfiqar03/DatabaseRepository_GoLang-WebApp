package dblibraryis

import (
	"database-repository/config"
	"database-repository/model/dbLibraryis"
)

func GetTransaction () ([]dblibraryis.Transaction, error, string) {
	var transaction []dblibraryis.Transaction

	nama := "Transaction Table"

	result := config.DBTransaction.
		Table("transactions").
		Preload("Transaction_Detail").
		Preload("Fine_Payment").
		Find(&transaction)

	return transaction, result.Error, nama
}