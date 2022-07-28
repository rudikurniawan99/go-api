package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("GORM_DB_HOST"),
		os.Getenv("GORM_DB_USER"),
		os.Getenv("GORM_DB_PASSWORD"),
		os.Getenv("GORM_DB_NAME"),
		os.Getenv("GORM_DB_PORT"),
		os.Getenv("GORM_DB_SSL_MODE"),
		os.Getenv("GORM_DB_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("==== DB ====")
		log.Fatal(err)
		log.Println("==== DB ====")
	} else {
		log.Println("==== DB ====")
		log.Println("= success to connect to db =")
		log.Println("==== DB ====")
	}

	return db
}
