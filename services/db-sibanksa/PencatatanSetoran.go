package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetPencatatanSetoran() ([]dbsibanksa.PencatatanSetoran, error, string) {
	var pencatatan_setoran []dbsibanksa.PencatatanSetoran

	nama := "Pencatatan Setoran Table"

	result := config.DBSiBanksa.
		Table("pencatatan_setoran").
		Preload("UserDetail").
		Preload("Jadwal_Pelaksanaan").
		Find(&pencatatan_setoran)

	return pencatatan_setoran, result.Error, nama
}
