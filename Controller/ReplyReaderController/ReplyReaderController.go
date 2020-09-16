package replyreadercontroller

import (
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

//ReplyListFC ...Sortのためのtype
type ReplyListFC []replyFC

//client用のreply
type replyFC struct {
	UserName  string
	Content   string
	LikeNum   int
	CreatedAt string
}

//GetList ...質問に対する読者の返信一覧
func GetList(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	replyList := reply.GetListByReader(questionID)
	replyListFC := convertReplyListForClient(replyList)

	sort.Sort(sort.Reverse(replyListFC))

	response := map[string][]replyFC{"Replies": replyListFC}
	return c.JSON(http.StatusOK, response)
}

//replyモデルリストをclient用のreplyリストに変換
func convertReplyListForClient(replyList []reply.Reply) ReplyListFC {
	replyListFC := []replyFC{}
	for _, reply := range replyList {
		user := user.Get(reply.UserID)
		rand.Seed(time.Now().UnixNano())

		rFC := replyFC{}
		rFC.UserName = user.Name
		rFC.Content = reply.Content
		rFC.LikeNum = rand.Intn(100)
		rFC.CreatedAt = reply.CreatedAt.Format("2006-01-02 15:04:05")

		replyListFC = append(replyListFC, rFC)
	}
	return replyListFC
}

//Interface
func (r ReplyListFC) Len() int {
	return len(r)
}

func (r ReplyListFC) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ReplyListFC) Less(i, j int) bool {
	return r[i].LikeNum < r[j].LikeNum
}
