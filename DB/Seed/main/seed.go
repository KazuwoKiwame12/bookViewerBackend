package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	book "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Book"
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
	genre "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Genre"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	userbook "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/UserBook"
)

func main() {
	//Userデータ
	u := user.User{}
	u.Name = "Author"
	u.Email = "hogehoge1@gmail.com"
	u.Password = "vsavnjdslkkn"
	u.Status = 1
	user.Create(u)
	u.Name = "Reader"
	u.Email = "hogehoge2@gmail.com"
	u.Password = "cdn;ikklnln"
	u.Status = 0
	user.Create(u)

	//genreデータ
	g := genre.Genre{}
	g.Name = "Python"
	genre.Create(g)

	//bookデータ
	b := book.Book{}
	b.GenreID = 1
	b.Title = "Python入門"
	b.Image = "hogehoge"
	b.Price = 2000
	b.AuthorID = 2
	b.Bio = "この本は初心者のための本です"
	book.Create(b)

	//Chapterデータ
	c := chapter.Chapter{}
	c.BookID = 1
	c.Number = 1
	chapter.Create(c)

	//Pageデータ
	p := page.Page{}
	p.ChapterID = 1
	p.Number = 1
	page.Create(p)
	p.Number = 2
	page.Create(p)
	p.Number = 3
	page.Create(p)

	//Sentenceデータ
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

	//Questionデータ
	q := question.Question{}
	q.UserID = 1
	q.SentenceID = 1
	q.Title = "Pythonについて"
	q.Content = "魅力を教えて～"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	q.CreatedAt = time.Now().In(jst)
	question.Create(q)

	//Replyデータ
	r := reply.Reply{}
	r.UserID = 1
	r.QuestionID = 1
	r.Content = "Pythonでは機械学習に特化したライブラリがたくさん存在します"
	r.CreatedAt = time.Now().In(jst)
	reply.Create(r)
	r.Content = "Pythonではデータサイエンスのためのグラフ系のライブラリが存在します"
	reply.Create(r)
	r.UserID = 2
	r.Content = "Pythonは配列操作が柔軟に行うことが出来ます"
	reply.Create(r)

	//UserBookデータ
	ub := userbook.UserBook{}
	ub.UserID = 1
	ub.BookID = 1
	ub.Status = 0
	ub.CreatedAt = time.Now().In(jst)
	userbook.Create(ub)

	//likeデータ
	//TODO
}
