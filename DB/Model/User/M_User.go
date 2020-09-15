package user

import db "github.com/KazuwoKiwame12/bookViewerBackend/DB"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Status   int
}

func Create(user User) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&user)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...ユーザーモデルの取得
func Get(id int) User {
	db := db.Connect()
	defer db.Close()

	user := User{}
	db.First(&user, id)

	return user
}

//GetList ...全てのユーザーモデルの取得
func GetList() []User {
	db := db.Connect()
	defer db.Close()

	userList := []User{}
	db.Find(&userList)

	return userList
}
