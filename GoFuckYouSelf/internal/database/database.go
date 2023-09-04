package internal

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=5432 sslmode-disable TimeZone=Asia/shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}, &Category{}, &Books{})
	
	return db;
}
