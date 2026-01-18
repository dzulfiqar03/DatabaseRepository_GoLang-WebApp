package dblibraryis

import (
	"database-repository/config"
	"database-repository/model/dbLibraryis"
)



func GetUsers() ([]dblibraryis.User, error, string) {
	var user []dblibraryis.User

	nama := "User Table"

	result := config.DBMember.
		Table("users").
		Preload("UserDetail").
		Preload("UserDetail.Roles").
		Preload("MemberDetails").
		Find(&user)	

	return user, result.Error, nama
}

