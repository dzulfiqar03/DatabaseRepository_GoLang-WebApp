package dbsibanksa

import (
	"database-repository/config"
	"database-repository/model/dbSibanksa"
)

func GetUserTransaction() ([]dbsibanksa.UserTransaction, error, string) {
	var user_transactions []dbsibanksa.UserTransaction

	nama := "User Transaction Table"

	result := config.DBSiBanksa.Table("user_transactions").
    Preload("UserDetail").     
    Preload("Banks").          
    Preload("PencatatanSetoran").
	Preload("PencatatanSetoran.Jadwal_Pelaksanaan").
    Find(&user_transactions)

	return user_transactions, result.Error, nama
}
