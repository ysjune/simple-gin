package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}

func generateSessionToken() string {
	// 랜덤한 16자리 문자열 토큰을 생성합니다.
	// 토큰생성의 안전한 방법은 아니므로, 운영환경에서는 사용하지 마세요.
	return strconv.FormatInt(rand.Int63(), 16)
}

func showLoginPage(c *gin.Context) {
	render(c, gin.H{"title": "Login"}, "login.html")
}

func performLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if isValidUser(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func logout(c *gin.Context) {
	// clear 쿠키
	c.SetCookie("token", "", -1, "", "", false, true)

	// 메인으로 redirect
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
