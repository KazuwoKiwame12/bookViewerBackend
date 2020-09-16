package ChapterController

import (
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type questionsData struct {
	QuestionID int
	Title      string
	UserName   string
	CreatedAt  time.Time
}
type QuestionsList struct {
	Questions []questionsData
}

// 質問一覧を取得
func GetQuestionList(c echo.Context) error {
	response := QuestionsList{}

	questionsDataList := []questionsData{}
	for _, q := range question.GetQuestionList() {
		user := user.Get(q.UserID)
		questionsData := questionsData{}
		questionsData.QuestionID = q.ID
		questionsData.UserName = user.Name
		questionsData.Title = q.Title
		questionsData.CreatedAt = q.Created_At

		questionsDataList = append(questionsDataList, questionsData)
	}
	response.Questions = questionsDataList
	return c.JSON(http.StatusOK, response)
}
