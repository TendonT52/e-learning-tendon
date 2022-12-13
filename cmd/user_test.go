package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/TendonT52/e-learning-tendon/internal/application"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestSignUp(t *testing.T) {
	apitest.New().
		Handler(router).
		Post("/api/v1/user/sign-up").
		JSON(`
		{
			"firstName":"firstNameTest",
			"lastName":"lastNameTest",
			"email":"username@email.com",
			"password":"passwordTest"
		}
		`).
		Expect(t).
		Status(http.StatusCreated).
		CookiePresent("token").
		Assert(jsonpath.
			Chain().
			Equal("firstName", "firstNameTest").
			Equal("lastName", "lastNameTest").
			Equal("email", "username@email.com").
			End()).
		End()
}

func TestSignUpValidationError(t *testing.T) {
	apitest.New().
		Handler(router).
		Post("/api/v1/user/sign-up").
		JSON(`
		{
			"lastName":"lastNameTest",
			"email":"emailTest",
			"password":"passwordTest"
		}
		`).
		Expect(t).
		Status(http.StatusBadRequest).
		Body(`
		{
			"message": "validation error",
			"error" :{
				"signUpReq.Email": "Email must be a valid email address",
        	    "signUpReq.FirstName": "FirstName is a required field"
			}
		}
		`).
		End()
}

func TestSignIn(t *testing.T) {
	apitest.New().
		Handler(router).
		Post("/api/v1/user/sign-up").
		JSON(`
		{
			"firstName":"firstNameTest",
			"lastName":"lastNameTest",
			"email":"username1@email.com",
			"password":"passwordTest"
		}
		`).
		Expect(t).
		Status(http.StatusCreated).
		CookiePresent("token").
		Assert(jsonpath.
			Chain().
			Equal("firstName", "firstNameTest").
			Equal("lastName", "lastNameTest").
			Equal("email", "username1@email.com").
			End()).
		End()
	apitest.New().
		Handler(router).
		Post("/api/v1/user/sign-in").
		JSON(`
		{
			"email":"username1@email.com",
			"password":"passwordTest"
		}
		`).
		Expect(t).
		Status(http.StatusOK).
		CookiePresent("token").
		Assert(jsonpath.
			Chain().
			Equal("firstName", "firstNameTest").
			Equal("lastName", "lastNameTest").
			Equal("email", "username1@email.com").
			End()).
		End()
}

func TestSignOut(t *testing.T) {
	_, token, err := application.UserServiceInstance.
		SignUp("firstNameTest2", "lastNameTest2", "username2@email.com", "passwordTest")
	if err != nil {
		t.Error("error while hook sign up")
	}
	req := fmt.Sprintf(`{ "refreshToken" : "%v" }`, token.Refresh)
	apitest.New().
		Handler(router).
		Post("/api/v1/auth/user/sign-out").
		Cookie("token", token.Access).
		JSON(req).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.
			Chain().
			Equal("message", "Sign out complete").
			End()).
		End()
}
