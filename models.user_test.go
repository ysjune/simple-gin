package main

import "testing"

func TestUserValidity(t *testing.T) {
	if !isValidUser("user1", "pass1") {
		t.Fail()
	}

	if isValidUser("user2", "pass1") {
		t.Fail()
	}

	if isValidUser("user1", "") {
		t.Fail()
	}

	if isValidUser("", "pass1") {
		t.Fail()
	}

	if isValidUser("User1", "pass1") {
		t.Fail()
	}
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	if isUsernameAvailable("user1") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}
