package main

import (
	"fmt"

	genre "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Genre"
)

func main() {
	model := genre.Genre{}
	model.Name = "Python"
	genre.Create(model)
	fmt.Println(genre.Get(1))
	fmt.Println(genre.GetList())
}
