package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetKepengurusan() ([]dbsibanksa.Kepengurusan, error, string) {
	var kepengurusan []dbsibanksa.Kepengurusan

	nama := "Kepengurusan Table"

	result := config.DBSiBanksa.
		Table("kepengurusans").
		Preload("UserDetail").
		Preload("UserDetail.RT").
		Preload("Gender").
		Find(&kepengurusan)

	return kepengurusan, result.Error, nama
}
