package dbnetabot

import (
	"database-repository/config"
	"database-repository/model/dbNetabot"
)

func GetChat() ([]dbnetabot.UserChat, error, string) {
	var userchat []dbnetabot.UserChat

	nama := "User Chat Table"

	result := config.DBNetabot.Table("user_chat").
		Preload("User").
		Preload("User.UserDetail").
		Find(&userchat)

	return userchat, result.Error, nama
}
