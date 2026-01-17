package userservice

import (
	"database-repository/config"
	"database-repository/model/User"

)

func GetUser() ([]User.User, error, string, []string) {
    var users []User.User
    
    // Simpan hasil query ke variabel 'result'
    result := config.DBLibraryIs.Find(&users)
	nama:="aku"
	dataList:=[]string{"ali"}
    
    // Kembalikan datanya dan error-nya (jika ada)
    return users, result.Error, nama, dataList
}