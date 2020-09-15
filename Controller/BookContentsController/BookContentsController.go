package bookcontentscontroller

import (
	"net/http"
	"strconv"

	bookservice "github.com/KazuwoKiwame12/bookViewerBackend/Service/BookService"
	"github.com/labstack/echo/v4"
)

//GetBookContents ...本の内容を取得
func GetBookContents(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))
	response := bookservice.GetBookContentsForClient(bookID)
	return c.JSON(http.StatusOK, response)
}
