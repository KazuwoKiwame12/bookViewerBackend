package main

import (
	"fmt"
	"net/http"
	"os"

	bookcontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/BookContentController"
	questioncontentcontroller "github.com/KazuwoKiwame12/bookViewerBackend/Controller/QuestionContentController"
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
	e.GET("/api/book/mine/:id", bookcontentcontroller.GetContent)
	e.GET("/api/question/:id/content", questioncontentcontroller.GetContent)
	/*
		e.Get("/api/chapter/:id", Controller当てはめる)
		e.Get("/api/question/:id/author/answer", Controller当てはめる)
		e.Get("/api/question/:id/reader/answer", Controller当てはめる)
		e.Get("/api/question/:id/page", Controller当てはめる)
		e.Get("/api/question/search/:title", Controller当てはめる)
		e.Get("/api/question/search/sentence/:id", Controller当てはめる)

		e.POST("/api/question/create", Controller当てはめる)
		e.POST("/api/question/reply", Controller当てはめる)
	*/

	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
