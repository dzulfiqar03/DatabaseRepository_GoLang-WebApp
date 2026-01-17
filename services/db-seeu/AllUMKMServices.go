package dbseeu

import (
	"database-repository/config"
	"database-repository/model/dbSeeu"
)

func GetAllUmkm() ([]dbseeu.AllUmkm, error, string) {
	var allUmkm []dbseeu.AllUmkm

	nama := "AllUmkm Table"

	result := config.DBSeeU.Table("allumkm").
		Preload("UMKMDetail").
		Find(&allUmkm)

	return allUmkm, result.Error, nama
}
