package SibanksaController

import (
	"database-repository/services/db-sibanksa"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title      string
	Rows       interface{} // Menampung data apapun (User, Roles, dll)
	Total      int
	TitleTable string
	IdTable    string
	IdForm     string
	TitleForm  string
}

func Index(c *gin.Context) {
	dataUser, errUser, namaUser := dbsibanksa.GetUsers()
	dataRoles, errRoles, namaRoles := dbsibanksa.GetRoles()
	dataRT, errRT, namaRT := dbsibanksa.GetRT()
	dataGender, errGender, namaGender := dbsibanksa.GetGender()
	dataBanks, errBanks, namaBanks := dbsibanksa.GetBanks()
	dataSampah, errSampah, namaSampah := dbsibanksa.GetSampah()
	dataDocument, errDocument, namaDocument := dbsibanksa.GetDocument()
	dataEvidence, errEvidence, namaEvidence := dbsibanksa.GetEvidence()
	dataJadwal, errJadwal, namaJadwal := dbsibanksa.GetJadwal()
	dataKepengurusan, errKepengurusan, namaKepengurusan := dbsibanksa.GetKepengurusan()
	dataUserTransaction, errUserTransaction, namaUserTransaction := dbsibanksa.GetUserTransaction()

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
	var dataSidebar = []string{
		"db-sibanksa",
		"db-givent",
		"db-seeU",
		"db-netabot",
		"db-libraryis",
	}

	sections := []TableSection{
		{
			Title:      "data-User",
			Rows:       dataUser,
			Total:      len(dataUser),
			TitleTable: namaUser,
			IdTable:    "tableUser",
			IdForm:     "formUser",
			TitleForm:  "Form Data User",
		},
		{
			Title:      "data-Roles",
			Rows:       dataRoles,
			Total:      len(dataRoles),
			TitleTable: namaRoles,
			IdTable:    "tableRoles",
			IdForm:     "formRoles",
			TitleForm:  "Form Data Roles",
		},
		{
			Title:      "data-Gender",
			Rows:       dataGender,
			Total:      len(dataGender),
			TitleTable: namaGender,
			IdTable:    "tableGender",
			IdForm:     "formGender",
			TitleForm:  "Form Data Gender",
		},
		{
			Title:      "data-Banks",
			Rows:       dataBanks,
			Total:      len(dataBanks),
			TitleTable: namaBanks,
			IdTable:    "tableBanks",
			IdForm:     "formBanks",
			TitleForm:  "Form Data Banks",
		},
		{
			Title:      "data-Document",
			Rows:       dataDocument,
			Total:      len(dataDocument),
			TitleTable: namaDocument,
			IdTable:    "tableDocument",
			IdForm:     "formDocument",
			TitleForm:  "Form Data Dokumen",
		},
		{
			Title:      "data-Evidence",
			Rows:       dataEvidence,
			Total:      len(dataEvidence),
			TitleTable: namaEvidence,
			IdTable:    "tableEvidence",
			IdForm:     "formEvidence",
			TitleForm:  "Form Data Evidence",
		},
		{
			Title:      "data-RT",
			Rows:       dataRT,
			Total:      len(dataRT),
			TitleTable: namaRT,
			IdTable:    "tableRT",
			IdForm:     "formRT",
			TitleForm:  "Form Data RT",
		},
		{
			Title:      "data-Sampah",
			Rows:       dataSampah,
			Total:      len(dataSampah),
			TitleTable: namaSampah,
			IdTable:    "tableSampah",
			IdForm:     "formSampah",
			TitleForm:  "Form Data Sampah",
		},
		{
			Title:      "data-Jadwal",
			Rows:       dataJadwal,
			Total:      len(dataJadwal),
			TitleTable: namaJadwal,
			IdTable:    "tableJadwal",
			IdForm:     "formJadwal",
			TitleForm:  "Form Data Jadwal",
		},
		{
			Title:      "data-Kepengurusan",
			Rows:       dataKepengurusan,
			Total:      len(dataKepengurusan),
			TitleTable: namaKepengurusan,
			IdTable:    "tableKepengurusan",
			IdForm:     "formKepengurusan",
			TitleForm:  "Form Data Kepengurusan",
		},
		{
			Title:      "data-UserTransaction",
			Rows:       dataUserTransaction,
			Total:      len(dataUserTransaction),
			TitleTable: namaUserTransaction,
			IdTable:    "tableUserTransaction",
			IdForm:     "formUserTransaction",
			TitleForm:  "Form Data User Transaction",
		},
	}

	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":           "db-sibanksa",
		"title":          "Halaman Database Sibanksa",
		"arrDataSidebar": dataSidebar,
		"sections":       sections,
	})

}
