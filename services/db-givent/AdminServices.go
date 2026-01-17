package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetAdmin() ([]dbgivent.Admin, error, string) {
	var admin []dbgivent.Admin

	nama := "Admin Table"

	result := config.DBGivent.Table("admin").
		Find(&admin)

	return admin, result.Error, nama
}
