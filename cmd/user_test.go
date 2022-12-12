package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestSignUp(t *testing.T) {
	apitest.New().
		Handler(setupRouter()).
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
		Status(http.StatusOK).
		Body(`{"param": "value"}`).
		End()
}

func TestSignUpFirstNameRequire(t *testing.T) {
	apitest.New().
		Handler(setupRouter()).
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
