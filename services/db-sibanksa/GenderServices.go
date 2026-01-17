package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetGender() ([]dbsibanksa.Gender, error, string) {
	var gender []dbsibanksa.Gender

	nama := "Gender Table"

	result := config.DBSiBanksa.
	Table("genders").
	Find(&gender)

	return gender, result.Error, nama
}
