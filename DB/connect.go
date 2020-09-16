package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//Connect ...DBとの接続
func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("env load error")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
