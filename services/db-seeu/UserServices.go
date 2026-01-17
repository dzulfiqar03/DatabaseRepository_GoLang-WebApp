package dbseeu

import (
	"database-repository/config"
	"database-repository/model/dbSeeu"
)

func GetUsers() ([]dbseeu.User, error, string) {
	var users []dbseeu.User

	nama := "User Table"

	result := config.DBSeeU.Table("users").
	Preload("AllUmkm").
	Preload("AllUmkm.Cities").
	Preload("AllUmkm.Category").
		Find(&users)

	return users, result.Error, nama
}
