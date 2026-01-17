package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetJadwal() ([]dbsibanksa.Jadwal_Pelaksanaan, error, string) {
	var jadwal_pelaksanaan []dbsibanksa.Jadwal_Pelaksanaan

	nama := "Jadwal Pelaksanaan Table"

	result := config.DBSiBanksa.
		Table("jadwal_pelaksanaan").
		Preload("UserDetail").
		Preload("UserDetail.RT").
		Find(&jadwal_pelaksanaan)

	return jadwal_pelaksanaan, result.Error, nama
}
