package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetRoles() ([]dbsibanksa.Roles, error, string) {
	var roles []dbsibanksa.Roles

	nama := "Roles Table"

	result := config.DBSiBanksa.
	Table("roles").
	Find(&roles)

	return roles, result.Error, nama
}
