package LibraryisController

import (
	"net/http"

	"database-repository/services/db-libraryis"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title      string
	Rows       interface{} // Menampung data apapun (UserDetail, Member_Details, dll)
	Total      int
	TitleTable string
	IdTable    string
	IdForm     string
	TitleForm  string
}

func Index(c *gin.Context) {
	dataUser, errUser, namaUser := dblibraryis.GetUsers()
	dataBook, errBook, namaBook := dblibraryis.GetBooks()
	dataTransaction, errTransaction, namaTransaction := dblibraryis.GetTransaction()
	if errUser != nil ||
		errBook != nil ||
		errTransaction != nil {
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
			IdTable:    "tableUserDetail",
			IdForm:     "formUserDetail",
			TitleForm:  "Form Data UserDetail",
		},
		{
			Title:      "data-Book",
			Rows:       dataBook,
			Total:      len(dataBook),
			TitleTable: namaBook,
			IdTable:    "tableBook",
			IdForm:     "formBook",
			TitleForm:  "Form Data Book",
		},
		{
			Title:      "data-Transaction",
			Rows:       dataTransaction,
			Total:      len(dataTransaction),
			TitleTable: namaTransaction,
			IdTable:    "tableTransaction",
			IdForm:     "formTransaction",
			TitleForm:  "Form Data Transaction",
		},
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"page":        "db-libraryis",
		"title":       "Database LibraryIS",
		"arrDataSidebar": dataSidebar,
		"sections":    sections,
	})
}