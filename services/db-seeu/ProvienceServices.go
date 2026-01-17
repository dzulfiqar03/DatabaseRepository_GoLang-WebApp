package dbseeu

import (
	"database-repository/config"
	"database-repository/model/dbSeeu"
)

func GetProvience() ([]dbseeu.Provience, error, string) {
	var provience []dbseeu.Provience

	nama := "Province Table"

	result := config.DBSeeU.Table("provinces").
	Preload("Cities").
		Find(&provience)

	return provience, result.Error, nama
}
