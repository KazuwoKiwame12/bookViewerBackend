package replyauthorcontroller

import (
	"net/http"
	"strconv"

	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	"github.com/labstack/echo/v4"
)

type replyFC struct {
	Content   string `json:"content"`
	LikeNum   int    `json:"likeNum"`
	CreatedAt string `json:"createdAt"`
}

//GetList ...質問に対する著者の返信一覧
func GetList(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	replyList := reply.GetListByAuthor(questionID)
	replyListFC := convertReplyListForClient(replyList)

	response := map[string][]replyFC{"answers": replyListFC}
	return c.JSON(http.StatusOK, response)
}

//replyモデルリストをclient用のreplyリストに変換
func convertReplyListForClient(replyList []reply.Reply) []replyFC {
	replyListFC := []replyFC{}
	for index, reply := range replyList {
		rFC := replyFC{}
		rFC.Content = reply.Content
		rFC.LikeNum = 40 - index
		rFC.CreatedAt = reply.CreatedAt.Format("2006-01-02 15:04:05")

		replyListFC = append(replyListFC, rFC)
	}
	return replyListFC
}
