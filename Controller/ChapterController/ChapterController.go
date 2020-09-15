package ChapterController

import (
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Questions struct {
	QuestionID int       `json:"question_id"`
	Title      string    `json:"title"`
	UserName   string    `json:"user_name"`
	CreatedAt  time.Time `json:"created_at"`
}

// 質問一覧を取得
// TODO: 章ごとに取得する
func GetQuestionList(c echo.Context) error {
	result := question.GetQuestionList()

	return c.JSON(http.StatusOK, result)
}
