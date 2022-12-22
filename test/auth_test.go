package test

import (
	"fmt"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/config"
	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	. "github.com/onsi/ginkgo/v2"

	// . "github.com/onsi/gomega"
	// . "github.com/onsi/gomega/gstruct"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

var _ = Describe("test user route", Ordered, func() {
	var t GinkgoTInterface

	BeforeAll(func() {
		config.LoadConfigTest()
		config.SetupInstance()
		handlers.SetupRouter()
		db.UserDBInstance.Clear()
		db.JwtDBInstance.Clear()
		db.CourseDBInstance.Clear()
		db.LessonDBInstance.Clear()
		db.NodeDBInstance.Clear()
		t = GinkgoT()
	})

	var token = core.Token{}
	Context("sign Up", func() {
		It("success", func() {
			res := apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/sign-up").
				JSON(`
				{
					"firstName": "firstnametest1",	
					"lastName": "lastnametest1",
					"email": "email@test.com",
					"password": "password1"
				}
				`).
				Expect(t).
				CookiePresent("refresh_token").
				Status(http.StatusCreated).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("firstName", "firstnametest1").
						Equal("lastName", "lastnametest1").
						Equal("email", "email@test.com").
						Present("accessToken").
						NotPresent("password").
						End(),
				).
				End().Response
			cookies := getCookie(res)
			body := getJSON(res)
			token = core.Token{
				Access:  body["accessToken"].(string),
				Refresh: cookies["refresh_token"],
			}
		})
	})

	Context("sign out", func() {
		It("success", func() {
			apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/sign-out").
				JSON(fmt.Sprintf(`{"accessToken": "%s"}`, token.Access)).
				Cookie("refresh_token", token.Refresh).
				Expect(t).
				CookieNotPresent("refresh_token").
				Status(http.StatusOK).
				End()
		})
	})

	Context("sign In", func() {
		It("success", func() {
			res := apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/sign-in").
				JSON(`
				{
					"email": "email@test.com",
					"password": "password1"
				}
				`).
				Expect(t).
				CookiePresent("refresh_token").
				Status(http.StatusOK).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("firstName", "firstnametest1").
						Equal("lastName", "lastnametest1").
						Equal("email", "email@test.com").
						NotPresent("password").
						Present("accessToken").
						End(),
				).
				End().Response
			cookies := getCookie(res)
			body := getJSON(res)
			token = core.Token{
				Access:  body["accessToken"].(string),
				Refresh: cookies["refresh_token"],
			}
		})
	})

	Context("refresh token", func() {
		It("success", func() {
			res := apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/refresh-token").
				Cookie("refresh_token", token.Refresh).
				Expect(t).
				CookiePresent("refresh_token").
				Status(http.StatusOK).
				Assert(
					jsonpath.Chain().
						Present("accessToken").
						End(),
				).
				End().Response
			cookies := getCookie(res)
			body := getJSON(res)
			token = core.Token{
				Access:  body["accessToken"].(string),
				Refresh: cookies["refresh_token"],
			}
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
