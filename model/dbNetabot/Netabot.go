package dbnetabot

type User struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Email      string     `json:"email"`
	UserDetail UserDetail `gorm:"foreignKey:UserID" json:"user_details"`
	UserChat   []UserChat `gorm:"foreignKey:UserID" json:"user_chat"`
}

type UserDetail struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	UserID   string `gorm:"column:id_user" json:"id_user"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Roles    string `json:"roles"`
}

type UserChat struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	UserID       string `gorm:"column:id_user" json:"id_user"`
	User         User   `gorm:"foreignKey:UserID"`
	Chat         string `json:"chat"`
	Bot_Response string `json:"bot_response"`
	Session_Key  string `json:"session_key"`
}

type Product struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Price       float64 `json:"price"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Rating      string `json:"rating"`
	Sold        string `json:"sold"`
}
