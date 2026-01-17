package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetDocument() ([]dbsibanksa.Document_Archivers, error, string) {
	var document_archivers []dbsibanksa.Document_Archivers

	nama := "Document Archive Table"
	
	result := config.DBSiBanksa.
		Table("document_archivers").
		Preload("UserDetail").
		Preload("UserDetail.RT").
		Find(&document_archivers)

	return document_archivers, result.Error, nama
}
