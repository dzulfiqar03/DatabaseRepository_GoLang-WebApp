package dblibraryis

import "time"

type User struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	Email         string         `json:"email"`
	UserDetail    UserDetail     `gorm:"foreignKey:UserID" json:"user_details"`
	MemberDetails Member_Details `gorm:"foreignKey:UserID" json:"member_details"`
}

type Roles struct {
	ID         int          `gorm:"primaryKey" json:"id"`
	Role       string       `json:"role"`
	UserDetail []UserDetail `gorm:"foreignKey:RolesID" json:"user_details"`
}

func (Roles) TableName() string {
	return "roles"
}

type UserDetail struct {
	ID               int    `gorm:"primaryKey" json:"id"`
	UserID           string `gorm:"column:id_user" json:"id_user"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Telephone_Number string `json:"telephone_number"`

	RolesID int   `gorm:"column:id_roles" json:"id_role"`
	Roles   Roles `gorm:"foreignKey:RolesID"`

	Status string `json:"status"`
}

func (UserDetail) TableName() string {
	return "user_details"
}

type Member_Details struct {
	ID                int     `gorm:"primaryKey" json:"id"`
	UserID            string  `gorm:"column:id_user" json:"id_userdetail"`
	Membership_Status string  `json:"membership_status"`
	Borrowing_Count   int     `json:"borrowing_count"`
	Total_Fine        float64 `json:"total_fine"`
}

func (Member_Details) TableName() string {
	return "member_details"
}

type Book struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Title       string      `json:"title"`
	Book_Detail Book_Detail `gorm:"foreignKey:BookID" json:"book_detail"`
}

func (Book) TableName() string {
	return "books"
}

type Book_Detail struct {
	ID               int    `gorm:"primaryKey" json:"id"`
	BookID           int    `gorm:"column:id_book" json:"id_book"`
	Author           string `json:"author"`
	Publisher        string `json:"publisher"`
	Year             int    `json:"year"`
	ISBN             string `json:"isbn"`
	Category         string `json:"category"`
	Language         string `json:"language"`
	Pages            int    `json:"pages"`
	Copies_Available int    `json:"quantity"`
}

func (Book_Detail) TableName() string {
	return "book_details"
}

type Transaction struct {
	ID               string    `gorm:"primaryKey;column:id" json:"id"`
	ID_Member        string    `gorm:"column:id_member" json:"id_member"` // Tetap teks, tapi mapping kolom benar
	Transaction_Date time.Time `gorm:"column:transaction_date" json:"transaction_date"`
	Due_Date         time.Time `gorm:"column:due_date" json:"due_date"`
	Return_Date      time.Time `gorm:"column:return_date" json:"return_date"`
	Status           string    `gorm:"column:status" json:"status"`
	Fine_Amount      float64   `gorm:"column:fine_amount" json:"fine_amount"`

	Transaction_Detail []Transaction_Detail `gorm:"foreignKey:TransactionID" json:"transaction_detail"`
	Fine_Payment       []Fine_Payment       `gorm:"foreignKey:TransactionID" json:"fine_payment"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type Transaction_Detail struct {
	ID            int     `gorm:"primaryKey" json:"id"`
	TransactionID string  `gorm:"column:id_transaction" json:"id_transaction"`
	BookID        int     ` json:"id_book"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

func (Transaction_Detail) TableName() string {
	return "transaction_details"
}

type Fine_Payment struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	TransactionID  string    `gorm:"column:id_transaction" json:"id_transaction"`
	ID_Member      string    ` json:"id_member"`
	Fine_Amount    float64   `json:"fine_amount"`
	Paid_Amount    float64   `json:"paid_amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	Paid_Date      time.Time `json:"paid_date"`
	Payment_Method string    `json:"payment_method"`
}
