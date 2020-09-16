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
// FIX: 章ごとに取得する
func GetQuestionList(c echo.Context) error {
	response := QuestionsList{}

	questionsDataList := []questionsData{}
	for _, question := range question.GetQuestionList() {
		user := user.Get(question.UserID)
		questionsData := questionsData{}
		questionsData.QuestionID = question.ID
		questionsData.UserName = user.Name
		questionsData.Title = question.Title
		questionsData.CreatedAt = question.Created_At

		questionsDataList = append(questionsDataList, questionsData)
	}
	response.Questions = questionsDataList
	return c.JSON(http.StatusOK, response)
}
