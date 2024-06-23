package schemas

import "time"

type SupplierSchema struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	NamaSupplier string    `json:"nama_supplier"`
	Alamat       string    `json:"alamat"`
	NoTelp       string    `json:"no_telp"`
}
