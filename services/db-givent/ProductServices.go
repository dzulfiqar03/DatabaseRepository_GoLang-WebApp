package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetProduct() ([]dbgivent.Product, error, string) {
	var product []dbgivent.Product

	nama := "Product Table"

	result := config.DBGivent.Table("product").
		Find(&product)

	return product, result.Error, nama
}
