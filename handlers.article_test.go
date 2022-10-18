package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ----- 홈페이지에 대한 GET 요청이 인증되지 않은 사용자에 대해 HTTP 200 코드가 있는 홈페이지를 반환하는지 테스트합니다 ----- //
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// ---- 위의 라우터로 보낼 요청을 생성합니다 ----- //
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// ----- HTTP status 코드가 200인지 확인합니다 ----- //
		statusOK := w.Code == http.StatusOK

		// ------ 페이지의 제목이 'Home Page' 인지 테스트합니다 ------- //
		// --- HTML 페이지를 파싱하고 처리할 수 있는 라이브러리를 사용하여 훨씬 디테일한 테스트를 수행합니다 --- //
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// --- Accept 가 json 일 때의 요청을 테스트합니다 --- //
func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article
		err = json.Unmarshal(p, &articles)

		return err == nil && len(articles) >= 2 && statusOK
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var a article
		err = xml.Unmarshal(p, &a)

		return err == nil && a.ID == 1 && len(a.Title) >= 0 && statusOK
	})

}
