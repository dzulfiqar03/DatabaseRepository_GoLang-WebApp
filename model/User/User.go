package User

type User struct {
    ID       string   `gorm:"primaryKey" json:"id"`
    Email    string `json:"email"`
}
