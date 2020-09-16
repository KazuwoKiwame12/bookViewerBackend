package ChapterController

import (
	"net/http"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type questionFC struct {
	QuestionID int
	Title      string
	UserName   string
	CreatedAt  string
}
type QuestionListFC struct {
	Questions []questionFC
}

//GetQuestionList 質問一覧を取得
func GetQuestionList(c echo.Context) error {
	questionList := question.GetQuestionList()
	questionListFC := convertQuestionListForClient(questionList)

	response := QuestionListFC{Questions: questionListFC}
	return c.JSON(http.StatusOK, response)
}

func convertQuestionListForClient(questionList []question.Question) []questionFC {
	questionListFC := []questionFC{}
	for _, q := range questionList {
		user := user.Get(q.UserID)
		qFC := questionFC{}
		qFC.QuestionID = q.ID
		qFC.UserName = user.Name
		qFC.Title = q.Title
		qFC.CreatedAt = q.CreatedAt.Format("2006-01-02 15:04:05")

		questionListFC = append(questionListFC, qFC)
	}
	return questionListFC
}
