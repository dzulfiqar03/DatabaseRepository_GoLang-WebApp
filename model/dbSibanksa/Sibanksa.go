package dbsibanksa

import "time"

type User struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Email      string     `json:"email"`
	UserDetail UserDetail `gorm:"foreignKey:UserID"`
}

type Roles struct {
	ID         int          `gorm:"primaryKey" json:"id"`
	Role       string       `json:"role"`
	UserDetail []UserDetail `gorm:"foreignKey:RolesID" json:"user_details"`
}

func (Roles) TableName() string {
	return "roles"
}

type RT struct {
	ID         int          `gorm:"primaryKey" json:"id"`
	RT         string       `json:"RT"`
	UserDetail []UserDetail `gorm:"foreignKey:RTID" json:"user_details"`
}

func (RT) TableName() string {
	return "rt_perumahan"
}

type Gender struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	Gender       string         `json:"gender"`
	UserDetail   []UserDetail   `gorm:"foreignKey:GenderID" json:"user_details"`
	Kepengurusan []Kepengurusan `gorm:"foreignKey:GenderID" json:"kepengurusans"`
}

func (Gender) TableName() string {
	return "genders"
}

type UserDetail struct {
	ID               int    `gorm:"primaryKey" json:"id"`
	UserID           string `gorm:"column:id_user" json:"id_user"`
	Username         string `json:"userName"`
	FullName         string `json:"fullName"`
	RTID             int    `gorm:"column:id_rt" json:"id_rt"`
	RT               RT     `gorm:"foreignKey:RTID;references:ID"`
	Address          string `json:"address"`
	Telephone_Number string `json:"telephone_number"`
	GenderID         int    `json:"id_gender"`
	Gender           Gender `gorm:"foreignKey:GenderID"`

	RolesID int   `gorm:"column:id_roles" json:"id_roles"`
	Roles   Roles `gorm:"foreignKey:RolesID"`

	Status             string               `json:"status"`
	Kepengurusan	   []Kepengurusan       `gorm:"foreignKey:UserDetailID" json:"kepengurusans"`
	Document_Archivers []Document_Archivers `gorm:"foreignKey:UserDetailID" json:"document_archivers"`
	Evidence_Archivers []Evidence_Archivers `gorm:"foreignKey:UserDetailID" json:"evidence_archivers"`
	Jadwal_Pelaksanaan []Jadwal_Pelaksanaan `gorm:"foreignKey:UserDetailID" json:"jadwal_pelaksanaan"`
	Sampah             []Sampah             `gorm:"foreignKey:UserDetailID" json:"sampah"`
	UserTransaction    []UserTransaction    `gorm:"foreignKey:UserDetailID" json:"user_transactions"`
}

type Document_Archivers struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	UserDetailID int        `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail   UserDetail `gorm:"foreignKey:UserDetailID"`
	Name         string     `json:"name"`
}

type Evidence_Archivers struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	UserDetailID int        `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail   UserDetail `gorm:"foreignKey:UserDetailID"`
	Name         string     `json:"name"`
}

type Jadwal_Pelaksanaan struct {
	ID                int                 `gorm:"primaryKey" json:"id"`
	UserDetailID      int                 `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail        UserDetail          `gorm:"foreignKey:UserDetailID"`
	Tanggal_Setoran   time.Time           `json:"tanggal_setoran"`
	PencatatanSetoran []PencatatanSetoran `gorm:"foreignKey:JadwalID" json:"pencatatan_setoran"`
}

type Kepengurusan struct {
	ID               int        `gorm:"primaryKey" json:"id"`
	UserDetailID     int        `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail       UserDetail `gorm:"foreignKey:UserDetailID"`
	Username         string     `json:"userName"`
	FullName         string     `json:"fullName"`
	Address          string     `json:"address"`
	Telephone_Number string     `json:"telephone_number"`
	GenderID         int        `gorm:"column:id_gender" json:"id_gender"`
	Gender           Gender     `gorm:"foreignKey:GenderID"`
	Divisi           string     `json:"divisi"`
}

type Sampah struct {
	ID                     int                      `gorm:"primaryKey" json:"id"`
	Nama_Sampah            string                   `json:"nama_sampah"`
	Harga                  int                      `json:"harga"`
	Satuan                 string                   `json:"satuan"`
	Kategori               string                   `json:"kategori"`
	UserDetailID           int                      `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail             UserDetail               `gorm:"foreignKey:UserDetailID" json:"user_detail"`
	PencatatanSetoranItems []PencatatanSetoranItems `gorm:"foreignKey:SampahID" json:"pencatatan_setoran_items"`
}

type PencatatanSetoran struct {
	ID                     int                      `gorm:"primaryKey" json:"id"`
	JadwalID               int                      `gorm:"column:id_jadwal" json:"id_jadwal"`
	Jadwal_Pelaksanaan     Jadwal_Pelaksanaan       `gorm:"foreignKey:JadwalID"`
	UserDetailID           int                      `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail             UserDetail               `gorm:"foreignKey:UserDetailID"`
	Tanggal_Setoran        time.Time                `json:"tanggal_setoran"`
	Total_Setoran          int                      `json:"total_setoran"`
	PencatatanSetoranItems []PencatatanSetoranItems `gorm:"foreignKey:PencatatanSetoranID" json:"pencatatan_setoran_items"`
	UserTransaction        []UserTransaction        `gorm:"foreignKey:PencatatanSetoranID" json:"user_transactions"`
}

type PencatatanSetoranItems struct {
	ID                  int               `gorm:"primaryKey" json:"id"`
	PencatatanSetoranID int               `gorm:"column:pencatatan_setoran_id" json:"pencatatan_setoran_id"`
	PencatatanSetoran   PencatatanSetoran `gorm:"foreignKey:pencatatan_setoran_id"`
	SampahID            int               `gorm:"column:sampah_id" json:"sampah_id"`
	Sampah              Sampah            `gorm:"foreignKey:SampahID"`
	Jumlah              int               `json:"jumlah"`
	Harga_Satuan        int               `json:"harga_satuan"`
	SubTotal            int               `json:"subtotal"`
}

type Banks struct {
	ID              int               `gorm:"primaryKey" json:"id"`
	Transfer_Code   string            `json:"transfer_code"`
	Name            string            `json:"name"`
	Short_Name      string            `json:"short_name"`
	Swift_Code      string            `json:"swift_code"`
	Logo            string            `json:"logo"`
	UserTransaction []UserTransaction `gorm:"foreignKey:BanksID" json:"user_transactions"`
}

type UserTransaction struct {
	ID                  int               `gorm:"primaryKey" json:"id"`
	UserDetailID        int               `gorm:"column:id_userdetail" json:"id_userdetail"`
	UserDetail          UserDetail        `gorm:"foreignKey:UserDetailID"`
	BanksID             int               `gorm:"column:id_bank" json:"id_bank"`
	Banks               Banks             `gorm:"foreignKey:BanksID"`
	Nomor_Rekening      string            `json:"nomor_rekening"`
	PencatatanSetoranID int               `gorm:"column:pencatatan_setoran_id" json:"pencatatan_setoran_id"`
	PencatatanSetoran   PencatatanSetoran `gorm:"foreignKey:PencatatanSetoranID"`
	Bukti_Pembayaran    string            `json:"bukti_pembayaran"`
}
