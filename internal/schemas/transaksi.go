package schemas

import "time"

type TransaksiSchema struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	TotalHarga       uint      `json:"total_harga"`
	MetodePembayaran string    `json:"metode_pembayaran"`
}
