package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetCustomer() ([]dbgivent.Customer, error, string) {
	var customer []dbgivent.Customer

	nama := "Customer Table"

	result := config.DBGivent.Table("customer").
		Find(&customer)

	return customer, result.Error, nama
}
