package SeeuController

import (
	"database-repository/services/db-seeu"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title      string
	Rows       interface{} // Menampung data apapun (User, Provience, dll)
	Total      int
	TitleTable string
	IdTable    string
	IdForm     string
	TitleForm  string
}

func Index(c *gin.Context) {
	dataUser, errUser, namaUser := dbseeu.GetUsers()
	dataProvience, errProvience, namaProvience := dbseeu.GetProvience()
	dataCities, errCities, namaCities := dbseeu.GetCities()
	dataCategory, errCategory, namaCategory := dbseeu.GetCategory()
	dataAllUmkm, errAllUmkm, namaAllUmkm := dbseeu.GetAllUmkm()

	if errUser != nil ||
		errProvience != nil ||
		errCategory != nil ||
		errCities != nil ||
		errAllUmkm != nil  {
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
			Title:      "data-Provience",
			Rows:       dataProvience,
			Total:      len(dataProvience),
			TitleTable: namaProvience,
			IdTable:    "tableProvience",
			IdForm:     "formProvience",
			TitleForm:  "Form Data Provience",
		},
		{
			Title:      "data-Category",
			Rows:       dataCategory,
			Total:      len(dataCategory),
			TitleTable: namaCategory,
			IdTable:    "tableCategory",
			IdForm:     "formCategory",
			TitleForm:  "Form Data Category",
		},
		{
			Title:      "data-Cities",
			Rows:       dataCities,
			Total:      len(dataCities),
			TitleTable: namaCities,
			IdTable:    "tableCities",
			IdForm:     "formCities",
			TitleForm:  "Form Data Cities",
		},
		{
			Title:      "data-AllUmkm",
			Rows:       dataAllUmkm,
			Total:      len(dataAllUmkm),
			TitleTable: namaAllUmkm,
			IdTable:    "tableAllUmkm",
			IdForm:     "formAllUmkm",
			TitleForm:  "Form Data AllUmkm",
		},
	}

	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":           "db-seeU",
		"title":          "Halaman Database SeeU",
		"arrDataSidebar": dataSidebar,
		"sections":       sections,
	})

}
