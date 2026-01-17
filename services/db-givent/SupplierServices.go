package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetSupplier() ([]dbgivent.Supplier, error, string) {
	var supplier []dbgivent.Supplier

	nama := "Supplier Table"

	result := config.DBGivent.Table("supplier").
		Find(&supplier)

	return supplier, result.Error, nama
}
