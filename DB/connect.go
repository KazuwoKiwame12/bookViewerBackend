package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	//書き換える!!しかし、commitしないこと!
	dsn := "自分のDATABASE_URL?sslmode=disable"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
