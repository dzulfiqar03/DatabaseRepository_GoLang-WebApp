package dbseeu

import (
	"database-repository/config"
	"database-repository/model/dbSeeu"
)

func GetCities() ([]dbseeu.Cities, error, string) {
	var cities []dbseeu.Cities

	nama := "City Table"

	result := config.DBSeeU.Table("cities").
		Find(&cities)

	return cities, result.Error, nama
}
