package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {

	// ----- Gin에서 기본 라우터 생성 ----- //
	router = gin.Default()

	// ------ 모든 템플릿 파일 로드 ------- //
	router.LoadHTMLGlob("templates/*")

	// ------ 경로 초기화하기 -------- //
	initializeRoutes()

	// ------ 어플리케이션 서버 구동 ------//
	router.Run()

}

// --- Request Header 의 'Accept' 에 따라 HTML, JSON, XML 로 렌더링합니다 --- //
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

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
