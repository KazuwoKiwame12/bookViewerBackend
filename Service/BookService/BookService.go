package bookservice

import (
	"fmt"

	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
)

//BookData ...本のデータ
type BookData struct {
	Book []chapterData
}

type chapterData struct {
	ID    int
	Pages []pageData
}

type pageData struct {
	ID        int
	Sentences []sentenceData
}

type sentenceData struct {
	ID          int
	Content     string
	HasQuestion bool
}

//GetBookContentsForClient ...本のデータをフロント向けに整形する
func GetBookContentsForClient(bookID int) BookData {
	response := BookData{}
	chapterDataList := []chapterData{}

	for _, chapter := range chapter.GetListByBook(bookID) {
		pageDataList := []pageData{}
		for _, page := range page.GetListByChapter(chapter.ID) {
			sentenceDataList := []sentenceData{}
			for _, sentence := range sentence.GetListByPage(page.ID) {
				data := sentenceData{ID: sentence.ID, Content: sentence.Content, HasQuestion: checkHasQuestion(sentence.ID)}
				sentenceDataList := append(sentenceDataList, data)
				fmt.Println(sentenceDataList)
			}
			pageDataList := append(pageDataList, pageData{ID: page.ID, Sentences: sentenceDataList})
			fmt.Println(pageDataList)
		}
		chapterDataList := append(chapterDataList, chapterData{ID: chapter.ID, Pages: pageDataList})
		fmt.Println(chapterDataList)
	}

	response.Book = chapterDataList
	return response
}

func checkHasQuestion(sentenceID int) bool {
	questionList := question.GetListBySentence(sentenceID)
	if len(questionList) == 0 {
		return false
	}

	return true
}
