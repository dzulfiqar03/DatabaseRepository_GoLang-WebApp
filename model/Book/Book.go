package Book


type Book struct {
    ID       string   `gorm:"primaryKey" json:"id"`
    Title    string `json:"title"`
}