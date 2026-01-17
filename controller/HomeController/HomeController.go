package HomeController

import (
	"database-repository/services/db-sibanksa"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title string
	Rows  interface{} // Menampung data apapun (User, Roles, dll)
	Total int
}

func Index(c *gin.Context) {

	dataUser, errUser, namaUser := dbsibanksa.GetUsers()
	dataRoles, errRoles,namaRoles := dbsibanksa.GetRoles()
	dataRT, errRT,namaRT := dbsibanksa.GetRT()
	dataGender, errGender ,namaGender:= dbsibanksa.GetGender()
	dataBanks, errBanks,namaBanks := dbsibanksa.GetBanks()
	dataSampah, errSampah,namaSampah := dbsibanksa.GetSampah()
	dataDocument, errDocument,namaDocument := dbsibanksa.GetDocument()
	dataEvidence, errEvidence,namaEvidence := dbsibanksa.GetEvidence()
	dataJadwal, errJadwal,namaJadwal := dbsibanksa.GetJadwal()
	dataKepengurusan, errKepengurusan,namaKepengurusan := dbsibanksa.GetKepengurusan()
	dataUserTransaction, errUserTransaction,namaUserTransaction := dbsibanksa.GetUserTransaction()

	if errUser != nil ||
		errRoles != nil ||
		errGender != nil ||
		errBanks != nil ||
		errDocument != nil ||
		errEvidence != nil ||
		errRT != nil ||
		errSampah != nil ||
		errJadwal != nil ||
		errKepengurusan != nil ||
		errUserTransaction != nil {
		// handle error
	}

	var data = []string{
		namaUser,
		namaRoles,
		namaGender,
		namaBanks,
		namaDocument,
		namaEvidence,
		namaRT,
		namaSampah,
		namaJadwal,
		namaKepengurusan,
		namaUserTransaction,
	}

	var dataSidebar = []string{
		"db-sibanksa",
		"db-givent",
		"db-seeU",
	}

	var dataTable = []string{
		namaUser,
		namaRoles,
		namaGender,
		namaBanks,
		namaDocument,
		namaEvidence,
		namaRT,
		namaSampah,
		namaJadwal,
		namaKepengurusan,
		namaUserTransaction,
	}

	

	sections := []TableSection{
		{Title: namaUser, Rows: dataUser, Total: len(dataUser)},
		{Title: namaRoles, Rows: dataRoles, Total: len(dataRoles)},
		{Title: namaGender, Rows: dataGender, Total: len(dataGender)},
		{Title: namaBanks, Rows: dataBanks, Total: len(dataBanks)},
		{Title: namaDocument, Rows: dataDocument, Total: len(dataDocument)},
		{Title: namaEvidence, Rows: dataEvidence, Total: len(dataEvidence)},
		{Title: namaRT, Rows: dataRT, Total: len(dataRT)},
		{Title: namaSampah, Rows: dataSampah, Total: len(dataSampah)},
		{Title: namaJadwal, Rows: dataJadwal, Total: len(dataJadwal)},
		{Title: namaKepengurusan, Rows: dataKepengurusan, Total: len(dataKepengurusan)},
		{Title: namaUserTransaction, Rows: dataUserTransaction, Total: len(dataUserTransaction)},
	}

	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":            "Dashboard",
		"title":           "Halaman Home",
		"sections":        sections,
		"arrData":         data,
		"dataTable":       dataTable,
		"arrDataSidebar":  dataSidebar,
	})

}
