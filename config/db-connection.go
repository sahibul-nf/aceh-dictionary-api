package config

import (
	"aceh-dictionary-api/bookmark"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/user"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")
	// dbPass := os.Getenv("DB_PASS")
	// dbPort := os.Getenv("DB_PORT")

	// Mysql
	dsn := "root:@tcp(127.0.0.1:3306)/acehnese_dictionary?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// Postgres
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&dictionary.Dictionary{}, &user.User{}, &bookmark.Bookmark{})
	fmt.Println("Database connected!")

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
