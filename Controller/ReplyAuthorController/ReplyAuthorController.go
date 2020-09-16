package replyauthorcontroller

import (
	"net/http"
	"strconv"
	"time"

	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	"github.com/labstack/echo/v4"
)

type replyData struct {
	Content   string
	LikeNum   int
	CreatedAt time.Time
}

//GetList ...質問に対する著者の返信一覧
func GetList(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	replyList := reply.GetListByAuthor(questionID)

	replyDataList := []replyData{}
	for index, reply := range replyList {
		replyData := replyData{}
		replyData.Content = reply.Content
		replyData.LikeNum = 40 - index
		replyData.CreatedAt = reply.CreatedAt

		replyDataList = append(replyDataList, replyData)
	}

	response := map[string][]replyData{"ReplyList": replyDataList}
	return c.JSON(http.StatusOK, response)
}
