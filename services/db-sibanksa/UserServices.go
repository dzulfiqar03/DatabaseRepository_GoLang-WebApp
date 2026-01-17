package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetUsers() ([]dbsibanksa.User, error, string) {
	var users []dbsibanksa.User

	nama := "User Table"

	result := config.DBSiBanksa.Table("users").Preload("UserDetail").Preload("UserDetail.RT").Preload("UserDetail.Gender").
		Find(&users)

	return users, result.Error, nama
}
