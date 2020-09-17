package main

import (
	"fmt"
	"net/http"
	"os"

	bookcontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/BookContentController"
	chaptercontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/ChapterController"
	questioncontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionContentController"
	questioncontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionController"
	questionsearchcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionSearchController"
	replyauthorcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/ReplyAuthorController"
	replycontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/ReplyController"
	replyreadercontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/ReplyReaderController"
	sentencecontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/SentenceController"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Print("env load error")
	}
	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.CORS())
	// REST API
	e.GET("/", helloWorld)
	e.POST("/api/question/create", questioncontroller.Post)
	e.POST("/api/question/reply", replycontroller.Post)
	e.GET("/api/book/mine/:id", bookcontentcontroller.GetContent)
	e.GET("/api/question/:id/content", questioncontentcontroller.GetContent)
	e.GET("/api/chapter/:id", chaptercontroller.GetQuestionList)
	e.GET("/api/question/:id/page", sentencecontroller.GetSentence)
	e.GET("/api/question/search/sentence/:id", questionsearchcontroller.GetListBySentence)
	e.GET("/api/question/search/:title", questionsearchcontroller.GetListByTitle)
	e.GET("/api/question/:id/author/answer", replyauthorcontroller.GetList)
	e.GET("/api/question/:id/reader/answer", replyreadercontroller.GetList)

	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
