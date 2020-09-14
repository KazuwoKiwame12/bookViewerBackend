package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Connect ...DBとの接続
func Connect() *gorm.DB {
	/*
		//.env取得できない
		err := godotenv.Load("github.com/KazuwoKiwame12/bookViewerBackend/.env")
		if err != nil {
			fmt.Print("env load error")
		}

		dsn := os.Getenv("DATABASE_URL")
	*/

	dsn := "DATABASE_URL?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
