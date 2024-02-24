package database

import (
	"fmt"
	"goauth/model"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDB() *gorm.DB {
	godotenv.Load()

	PORT := os.Getenv("DBPORT")

	port, _ := strconv.ParseUint(PORT, 10, 64)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"),port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.AutoMigrate(&model.Users{})

	return db
}

