package NetabotController

import (
	"database-repository/services/db-netabot"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title      string
	Rows       interface{} // Menampung data apapun (User, Chat, dll)
	Total      int
	TitleTable string
	IdTable    string
	IdForm     string
	TitleForm  string
}

func Index(c *gin.Context) {
	dataUser, errUser, namaUser := dbnetabot.GetUsers()
	dataChat, errChat, namaChat := dbnetabot.GetChat()
	dataProduct, errProduct, namaProduct := dbnetabot.GetProduct()
	if errUser != nil ||
		errChat != nil ||
		errProduct != nil {
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
			Title:      "data-Chat",
			Rows:       dataChat,
			Total:      len(dataChat),
			TitleTable: namaChat,
			IdTable:    "tableChat",
			IdForm:     "formChat",
			TitleForm:  "Form Data Chat",
		},

		{
			Title:      "data-Product",
			Rows:       dataProduct,
			Total:      len(dataProduct),
			TitleTable: namaProduct,
			IdTable:    "tableProduct",
			IdForm:     "formProduct",
			TitleForm:  "Form Data Product",
		},
	}

	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":           "db-netabot",
		"title":          "Halaman Database Netabot",
		"arrDataSidebar": dataSidebar,
		"sections":       sections,
	})

}
