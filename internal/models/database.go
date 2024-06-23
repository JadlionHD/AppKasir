package models

import (
	"errors"

	"github.com/JadlionHD/AppKasir/internal/constants"
	"github.com/JadlionHD/AppKasir/internal/schemas"
	"github.com/JadlionHD/AppKasir/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = database

	if err = DB.AutoMigrate(&schemas.UserSchema{},
		&schemas.TransaksiSchema{},
		&schemas.SupplierSchema{},
		&schemas.ProdukSchema{},
		&schemas.DetailTransaksiSchema{}); err == nil && DB.Migrator().HasTable(&schemas.UserSchema{}) {
		if err := DB.First(&schemas.UserSchema{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seed()
		}
	}
}

func seed() {
	var flags uint
	flags |= constants.PERM_ADMIN

	hashedPassword := utils.HashingPassword("admin")
	DB.Create(&schemas.UserSchema{Username: "admin", Password: hashedPassword, Permissions: flags})

}
