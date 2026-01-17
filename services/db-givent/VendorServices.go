package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetVendor() ([]dbgivent.Vendor, error, string) {
	var vendor []dbgivent.Vendor

	nama := "Vendor Table"

	result := config.DBGivent.Table("vendor").
		Find(&vendor)

	return vendor, result.Error, nama
}
