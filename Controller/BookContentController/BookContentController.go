package bookcontentcontroller

import (
	"net/http"
	"strconv"

	bookservice "github.com/KazuwoKiwame12/bookViewerBackend/Service/BookService"
	"github.com/labstack/echo/v4"
)

//GetContent ...本の内容を取得
func GetContent(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))
	response := bookservice.GetContentForClient(bookID)
	return c.JSON(http.StatusOK, response)
}
