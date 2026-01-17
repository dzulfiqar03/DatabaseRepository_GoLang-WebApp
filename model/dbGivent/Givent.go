package dbgivent

import (
	"time"
)

type Customer struct {
	ID                int    `gorm:"primaryKey" json:"id"`
	Customer_Name     string `json:"customer_name"`
	Customer_Username string `json:"cust_username"`
	Email             string `json:"email"`
	Phone_Number      string `json:"phone_number"`
}

type Admin struct {
	Email string `json:"email"`
}

type Product struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Product_Name string `json:"product_name"`
	Event        string `json:"event"`
	Price        int    `json:"price"`
	Status_Stock string `json:"status_stock"`
}

type Event struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Event_Name string `json:"event_name"`
	Price      int    `json:"price"`
}

type TransactionProduct struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Customer_Name string `json:"customer_name"`
	Product       string `json:"product"`
	Event         string `json:"event"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
	Total_Price   int    `json:"total_price"`
}

type Supplier struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Product_Type string    `json:"customer_name"`
	Quantity     int       `json:"quantity"`
	Price        int       `json:"price"`
	Total_Price  int       `json:"total_price"`
	Date         time.Time `json:"date"`
	Status_Order string    `json:"status_order"`
}

type Vendor struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Product_Type string    `json:"customer_name"`
	Quantity     int       `json:"quantity"`
	Price        int       `json:"price"`
	Total_Price  int       `json:"total_price"`
	Date         time.Time `json:"date"`
	Status_Order string    `json:"status_order"`
}
