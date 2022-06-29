package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectBd() *gorm.DB {
	var (
		DB  *gorm.DB
		err error
	)
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error in Bd")
	}
	return DB
}

func CloseConnectionBd(DB *gorm.DB) {
	bdC, _ := DB.DB()
	bdC.Close()
}
