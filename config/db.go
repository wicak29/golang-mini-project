package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectionDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("penyimpanan.db"), &gorm.Config{})
	return db, err
	// if err != nil {
	// 	panic(err.Error())
	// }
}
