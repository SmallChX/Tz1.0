package setting

import (
	"fmt"
	"jobfair2024/model"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := "host=localhost port=5432 dbname=jobfair2024 user=postgres password=" + dbPassword + " sslmode=prefer connect_timeout=10"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}

}

func MigrateDB() {
	db.AutoMigrate(&model.CompanyInformation{}, &model.Booth{}, &model.BoothRequest{}, &model.Notification{})
	err := db.AutoMigrate(
		&model.Booth{}, &model.UserAccount{}, &model.BoothRequest{},
		&model.AdminInformation{}, &model.StudentInformation{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	initBoothData()
	initAdminAccount()
}

func initBoothData() {
	var count int64
	GetDB().Model(&model.Booth{}).Count(&count)

	// Kiểm tra nếu dữ liệu đã được khởi tạo
	if count == 0 {
		// Dữ liệu chưa được khởi tạo, tiến hành khởi tạo
		prices := []int{15000000, 12000000} // Giả định có 2 mức giá

		// Tạo các bản ghi cho bảng Booth
		for i := 1; i <= 76; i++ {
			price := prices[0] // Mặc định sử dụng giá đầu tiên
			if i > 14 && i <= 28 || i > 36 && i <= 40 || i > 48 && i <= 64 || i > 68 && i <= 76 {
				price = prices[1] // Sử dụng giá thứ hai cho các ID này
			}

			booth := model.Booth{
				ID:    int64(i),
				Price: price,
			}

			err := GetDB().Create(&booth).Error
			if err != nil {
				log.Printf("Failed to create booth with ID %d: %v", i, err)
			}
		}

		fmt.Println("Booth data initialized successfully.")
	} else {
		fmt.Println("Booth data already initialized.")
	}
}

func initAdminAccount() {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("cse@admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("cant hash password")
	}
	var count int64
	GetDB().Model(&model.UserAccount{}).Count(&count)
	if count == 0 {
		userAccount := &model.UserAccount{
			Username:   "admin",
			Password:   string(hashPassword),
			Role:       "admin",
			Email:      nil,
			FirstLogin: false,
		}

		err = GetDB().Create(&userAccount).Error
		if err != nil {
			log.Printf("fail to create admin account")
		}
	}
}
