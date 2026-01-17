package dbseeu

type User struct {
    ID       string    `gorm:"primaryKey" json:"id"`
    FullName string    `json:"fullName"`
    Username string    `json:"userName"`
    Email    string    `json:"email"`
    Roles    string    `json:"role"`
    AllUmkm  AllUmkm `gorm:"foreignKey:UserID" json:"all_umkm"` // Gunakan slice [] jika 1 user bisa punya banyak UMKM
}

type Provience struct {
    ID     int      `gorm:"primaryKey" json:"id"`
    Code   string   `json:"code"`
    Name   string   `json:"name"`
    Cities Cities `gorm:"foreignKey:ProvienceID" json:"cities"` // Slice karena 1 provinsi banyak kota
}

func (Provience) TableName() string { return "provinces" }

type Cities struct {
    ID          int       `gorm:"primaryKey" json:"id"`
    ProvienceID int       `gorm:"column:province_id" json:"provience_id"`
    Name        string    `json:"name"`
    AllUmkm     []AllUmkm `gorm:"foreignKey:CitiesID" json:"allumkm"`
}

type Category struct {
    ID          int       `gorm:"primaryKey" json:"id"`
    Code        string    `json:"code"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    AllUmkm     []AllUmkm `gorm:"foreignKey:CategoryID" json:"allumkm"`
}


func (Category) TableName() string { return "category" }

type UMKMDetail struct {
    ID               int    `gorm:"primaryKey" json:"id"`
    AllUmkmID        int    `gorm:"column:umkm_id" json:"umkm_id"` // PERBAIKAN: Tambah gorm:
    Description      string `json:"description"`
    Email            string `json:"email"`
    Address          string `json:"address"`
    Telephone_Number string `json:"telephone_number"`
}

func (UMKMDetail) TableName() string { return "detailUmkm" }

type AllUmkm struct {
    ID         int        `gorm:"primaryKey" json:"id"`
    UserID     string     `gorm:"column:id_user" json:"id_user"`
    UMKM       string     `json:"umkm"`
    CitiesID   int        `gorm:"column:city_id" json:"city_id"`
    CategoryID int        `gorm:"column:category_id" json:"category_id"`
    
    // Tambahkan struct penampung agar bisa di-Preload
    Cities     Cities     `gorm:"foreignKey:CitiesID;references:ID" json:"city"`
    Category   Category   `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
    UMKMDetail UMKMDetail `gorm:"foreignKey:AllUmkmID" json:"detailUmkm"`
}

func (AllUmkm) TableName() string { return "allumkm" }