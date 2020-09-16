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

//ReplyDataList ...Sortのためのtype
type ReplyDataList []replyData

type replyData struct {
	UserName  string
	Content   string
	LikeNum   int
	CreatedAt time.Time
}

//GetList ...質問に対する読者の返信一覧
func GetList(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	replyList := reply.GetListByReader(questionID)

	replyDataList := []replyData{}
	for _, reply := range replyList {
		user := user.Get(reply.UserID)
		rand.Seed(time.Now().UnixNano())

		replyData := replyData{}
		replyData.UserName = user.Name
		replyData.Content = reply.Content
		replyData.LikeNum = rand.Intn(100)
		replyData.CreatedAt = reply.CreatedAt

		replyDataList = append(replyDataList, replyData)
	}

	var sortData ReplyDataList
	sortData = replyDataList
	sort.Sort(sortData)

	response := map[string][]replyData{"ReplyList": sortData}
	return c.JSON(http.StatusOK, response)
}

func (r ReplyDataList) Len() int {
	return len(r)
}

func (r ReplyDataList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ReplyDataList) Less(i, j int) bool {
	return r[i].LikeNum < r[j].LikeNum
}
