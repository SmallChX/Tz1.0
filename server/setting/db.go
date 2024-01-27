package setting

import (
	"jobfair2024/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	dsn := "host=localhost port=5432 dbname=jobfair2023 user=postgres password=123456 sslmode=prefer connect_timeout=10"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	
}

func MigrateDB() {
	err := db.AutoMigrate(&model.Booth{}, &model.UserAccount{}, &model.BoothRequest{})
	if err != nil {
		panic("failed to migrate database")
	}
}
