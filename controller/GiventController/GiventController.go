package GiventController

import (
	"database-repository/services/db-givent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableSection struct {
	Title      string
	Rows       interface{} // Menampung data apapun (Customer, Admin, dll)
	Total      int
	TitleTable string
	IdTable    string
	IdForm     string
	TitleForm  string
}

func Index(c *gin.Context) {
	dataCustomer, errCustomer, namaCustomer := dbgivent.GetCustomer()
	dataAdmin, errAdmin, namaAdmin := dbgivent.GetAdmin()
	dataProduct, errProduct, namaProduct := dbgivent.GetProduct()
	dataEvent, errEvent, namaEvent := dbgivent.GetEvent()
	dataTransactionProduct, errTransactionProduct, namaTransactionProduct := dbgivent.GetTransactionProduct()
	dataSupplier, errSupplier, namaSupplier := dbgivent.GetSupplier()
	dataVendor, errVendor, namaVendor := dbgivent.GetVendor()

	if errCustomer != nil ||
		errAdmin != nil ||
		errEvent != nil ||
		errProduct != nil ||
		errTransactionProduct != nil ||
		errSupplier != nil ||
		errVendor != nil {
		// handle error
	}

	var dataSidebar = []string{
		"db-sibanksa",
		"db-givent",
		"db-seeU",
	}

	sections := []TableSection{
		{
			Title:      "data-Customer",
			Rows:       dataCustomer,
			Total:      len(dataCustomer),
			TitleTable: namaCustomer,
			IdTable:    "tableCustomer",
			IdForm:     "formCustomer",
			TitleForm:  "Form Data Customer",
		},
		{
			Title:      "data-Admin",
			Rows:       dataAdmin,
			Total:      len(dataAdmin),
			TitleTable: namaAdmin,
			IdTable:    "tableAdmin",
			IdForm:     "formAdmin",
			TitleForm:  "Form Data Admin",
		},
		{
			Title:      "data-Event",
			Rows:       dataEvent,
			Total:      len(dataEvent),
			TitleTable: namaEvent,
			IdTable:    "tableEvent",
			IdForm:     "formEvent",
			TitleForm:  "Form Data Event",
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
		{
			Title:      "data-TransactionProduct",
			Rows:       dataTransactionProduct,
			Total:      len(dataTransactionProduct),
			TitleTable: namaTransactionProduct,
			IdTable:    "tableTransactionProduct",
			IdForm:     "formTransactionProduct",
			TitleForm:  "Form Data TransactionProduct",
		},

		{
			Title:      "data-Supplier",
			Rows:       dataSupplier,
			Total:      len(dataSupplier),
			TitleTable: namaSupplier,
			IdTable:    "tableSupplier",
			IdForm:     "formSupplier",
			TitleForm:  "Form Data Supplier",
		},

		{
			Title:      "data-Vendor",
			Rows:       dataVendor,
			Total:      len(dataVendor),
			TitleTable: namaVendor,
			IdTable:    "tableVendor",
			IdForm:     "formVendor",
			TitleForm:  "Form Data Vendor",
		},
	}

	// Langsung tampilkan HTML
	c.HTML(http.StatusOK, "layout", gin.H{
		"page":           "db-givent",
		"title":          "Halaman Database Givent",
		"arrDataSidebar": dataSidebar,
		"sections":       sections,
	})

}
