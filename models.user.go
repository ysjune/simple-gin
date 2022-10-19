package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

/*
이번 데모에서는 미리 정의된 유저 목록을 메모리에 저장합니다.
실제 어플리케이션에서는 DB로 관리합니다. 그리고 당연하게도 패스워드는
암호화하여 저장해야 합니다.
*/
var userList = []user{
	user{Username: "user1", Password: "pass1"},
	user{Username: "user2", Password: "pass2"},
	user{Username: "user3", Password: "pass3"},
}

// 유저명과 패스워드 조합이 유효한지 확인
func isValidUser(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("패스워드는 공백일 수 없습니다.")
	}
	if !isUsernameAvailable(username) {
		return nil, errors.New("이미 존재하는 유저입니다.")
	}

	u := user{Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil
}

// 유저명이 가능한지 체크
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
