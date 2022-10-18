package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// ----- 모든 기사 목록을 반환합니다 ----- //
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("기사를 찾을 수 없습니다")
}

// --- 새로운 기사를 만듭니다. --- //
func createNewArticle(title, content string) (*article, error) {
	// --- 기존 기사들 수에 +1 된 값을 아이디에 세팅합니다. --- //
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	// --- 기사 목록에 기사를 추가합니다. --- //
	articleList = append(articleList, a)

	return &a, nil
}
