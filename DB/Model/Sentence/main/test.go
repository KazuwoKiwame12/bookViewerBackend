package main

import (
	"bufio"
	"fmt"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
	"os"
)

func main() {
	files := [3]string{"FirstPage.txt", "SecondPage.txt", "ThirdPage.txt"}
	for index, file := range files {
		fp, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			model := sentence.Sentence{}
			model.PageID = index + 1
			model.Content = scanner.Text()
			sentence.Create(model)
			fmt.Println(scanner.Text())
		}
	}

	fmt.Println(sentence.Get(1))
	fmt.Println(sentence.GetListByPage(1))
}
