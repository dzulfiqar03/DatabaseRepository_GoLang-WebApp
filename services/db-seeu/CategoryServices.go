package dbseeu

import (
	"database-repository/config"
	"database-repository/model/dbSeeu"
)

func GetCategory() ([]dbseeu.Category, error, string) {
	var category []dbseeu.Category

	nama := "Category Table"

	result := config.DBSeeU.Table("category").
		Find(&category)

	return category, result.Error, nama
}
