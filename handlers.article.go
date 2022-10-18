package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	//// ---- Context의 HTML 메소드를 호출하여 템플릿을 렌더링합니다 ---- //
	//c.HTML(
	//	// ----- HTTP Status를 200(OK)로 설정합니다 ------ //
	//	http.StatusOK,
	//	// ----- index.html 템플릿을 사용합니다 ------ //
	//	"index.html",
	//	// ---- 페이지에서 사용하는 데이터를 전달합니다 ----- //
	//	gin.H{
	//		"title":   "Home Page",
	//		"payload": articles,
	//	},
	//)

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")

}

func getArticle(c *gin.Context) {
	// ----- 기사 ID가 유효한지 확인합니다 ----- //
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// ----- 기사가 존재하는지 확인합니다 ----- //
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")

		} else {
			// ---- 기사를 찾을 수 없는 경우 오류와 함께 중단합니다 ---- //
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// ---- URL에 잘못된 기사 ID가 지정된 경우 오류와 함께 중단합니다 ---- //
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// --- Request Header 의 'Accept' 에 따라 HTML, JSON, XML 로 렌더링합니다 --- //
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		// --- 기사 생성이 성공했을 경우 success 메시지 보여줍니다 --- //
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// --- 기사 생성중 에러가 발생할 경우 --- //
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
