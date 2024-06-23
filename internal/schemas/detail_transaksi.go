package schemas

import "time"

type DetailTransaksiSchema struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IDTransaksi uint      `json:"id_transaksi"`
	IDProduk    uint      `json:"id_produk"`
	Jumlah      int       `json:"jumlah"`
	HargaSatuan int       `json:"harga_satuan"`
}
