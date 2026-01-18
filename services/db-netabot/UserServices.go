package dbnetabot

import (
	"database-repository/config"
	"database-repository/model/dbNetabot"
)

func GetUsers() ([]dbnetabot.User, error, string) {
	var users []dbnetabot.User

	nama := "User Table"

	result := config.DBNetabot.Table("users").
	Preload("UserDetail").
	Preload("UserChat").
		Find(&users)

	return users, result.Error, nama
}
