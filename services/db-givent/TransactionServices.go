package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetTransactionProduct() ([]dbgivent.TransactionProduct, error, string) {
	var transaction []dbgivent.TransactionProduct

	nama := "Transaction Product Table"

	result := config.DBGivent.Table("transactionproduct").
		Find(&transaction)

	return transaction, result.Error, nama
}
