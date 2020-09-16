package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KazuwoKiwame12/bookViewerBackend/Controller/ChapterController"

	bookcontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/BookContentController"
	questioncontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionContentController"
	questioncontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionController"
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
	e.GET("/api/chapter/:id", ChapterController.GetQuestionList)
	e.GET("/api/question/:id/page", sentencecontroller.GetSentence)
	/*
		e.Get("/api/question/search/:title", Controller当てはめる)
		e.Get("/api/question/search/sentence/:id", Controller当てはめる)
	*/
	e.GET("/api/question/:id/author/answer", replyauthorcontroller.GetList)
	e.GET("/api/question/:id/reader/answer", replyreadercontroller.GetList)

	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
