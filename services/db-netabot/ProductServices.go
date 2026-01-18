package dbnetabot

import (
	"database-repository/config"
	"database-repository/model/dbNetabot"
)

func GetProduct() ([]dbnetabot.Product, error, string) {
	var product []dbnetabot.Product

	nama := "User Product Table"

	result := config.DBNetabot.Table("products").
		Find(&product)

	return product, result.Error, nama
}
