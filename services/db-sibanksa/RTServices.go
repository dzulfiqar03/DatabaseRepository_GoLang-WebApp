package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetRT() ([]dbsibanksa.RT, error, string) {
	var rt []dbsibanksa.RT

	nama := "RT Table"

	result := config.DBSiBanksa.
		Table("rt_perumahan").
		Find(&rt)

	return rt, result.Error, nama
}
