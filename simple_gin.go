package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// ----- Gin에서 기본 라우터 생성 ----- //
	router = gin.Default()

	// ------ 모든 템플릿 파일 로드 ------- //
	router.LoadHTMLGlob("templates/*")

	//router.GET("/", func(c *gin.Context) {
	//	// ------ Context의 HTML 메소드를 호출하여 템플릿을 렌더링합니다. ----- //
	//	c.HTML(
	//		// ----- HTTP 상태를 200(OK)에 세팅합니다 ------ //
	//		http.StatusOK,
	//		// ------ index.html 템플릿을 사용합니다 ------- //
	//		"index.html",
	//		// ----- 페이지에서 사용하는 데이터 전달 ------- //
	//		gin.H{
	//			"title": "Home Page",
	//		},
	//	)
	//})

	// ------ 경로 초기화하기 -------- //
	initializeRoutes()

	// ------ 어플리케이션 서버 구동 ------//
	router.Run()

}
