package main

func initializeRoutes() {

	// 인덱스 라우터 처리(Handle)
	router.GET("/", showIndexPage)

	router.GET("/u/register", showRegistrationPage)

	router.POST("/u/register", register)

	articleRoutes := router.Group("/article")
	{
		// /article/view/some_article_id 부분에 대한 GET 요청 처리
		articleRoutes.GET("/view/:article_id", getArticle)

		articleRoutes.GET("/create", showArticleCreationPage)

		articleRoutes.POST("/create", createArticle)
	}

}
