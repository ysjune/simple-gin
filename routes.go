package main

func initializeRoutes() {

	// 인덱스 라우터 처리(Handle)
	router.GET("/", showIndexPage)

	// /article/view/some_article_id 부분에 대한 GET 요청 처리
	router.GET("/article/view/:article_id", getArticle)
}
