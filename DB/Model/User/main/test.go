package main

import (
	"fmt"

	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
)

func main() {
	// model := user.User{}
	// model.Name = "Guest1"
	// model.Email = "hogehoge1@gmail.com"
	// model.Password = "vsavnjdslkkn"
	// model.Status = 1
	// user.Create(model)
	// model.Name = "Guest2"
	// model.Email = "hogehoge2@gmail.com"
	// model.Password = "cdn;ikklnln"
	// user.Create(model)

	fmt.Println(user.Get(1))
	fmt.Println(user.GetList())
}
