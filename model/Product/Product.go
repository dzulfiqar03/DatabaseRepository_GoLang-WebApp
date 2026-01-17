package Product

type Product struct {
    ID       int   `gorm:"primaryKey" json:"id"`
    Product_Name    string `json:"product_name"`
	Event    string `json:"event"`
	Price	 int `json:"price"`
	Status_Stock	 string `json:"status_stock"`
}


