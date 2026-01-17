package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetBanks() ([]dbsibanksa.Banks, error, string) {
	var bank []dbsibanksa.Banks

	nama := "Banks Table"

	result := config.DBSiBanksa.
	Table("banks").
	Find(&bank)

	return bank, result.Error, nama
}
