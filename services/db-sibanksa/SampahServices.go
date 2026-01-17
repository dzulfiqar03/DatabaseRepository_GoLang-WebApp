package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetSampah() ([]dbsibanksa.Sampah, error, string) {
	var sampah []dbsibanksa.Sampah

	nama := "Sampah Table"

	result := config.DBSiBanksa.
		Table("sampah").
		Preload("UserDetail").
		Preload("UserDetail.RT").
		Find(&sampah)

	return sampah, result.Error, nama
}
