package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetEvidence() ([]dbsibanksa.Evidence_Archivers, error, string) {
	var evidence_archivers []dbsibanksa.Evidence_Archivers

	nama := "Evidence Archive Table"
	
	result := config.DBSiBanksa.
		Table("evidence_archivers").
		Preload("UserDetail").
		Preload("UserDetail.RT").
		Find(&evidence_archivers)

	return evidence_archivers, result.Error, nama
}
