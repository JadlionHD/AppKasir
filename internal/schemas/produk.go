package schemas

import "time"

type ProdukSchema struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Stok      int       `json:"stok"`
	Harga     uint      `json:"harga"`
	Kategori  string    `json:"kategori"`
}
