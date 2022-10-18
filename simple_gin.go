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

	// ------ 경로 초기화하기 -------- //
	initializeRoutes()

	// ------ 어플리케이션 서버 구동 ------//
	router.Run()

}
