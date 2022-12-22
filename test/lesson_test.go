package test

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/config"
	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/auth"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	// . "github.com/onsi/gomega/gstruct"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

var _ = Describe("test lesson route", Ordered, func() {
	var t GinkgoTInterface

	var userAdmin = core.User{
		FirstName:      "admin",
		LastName:       "admin",
		Email:          "admin@test.com",
		HashedPassword: auth.HashPassword("12345678"),
		Role:           core.Admin,
	}
	var userTeacher = core.User{
		FirstName:      "teacher",
		LastName:       "teacher",
		Email:          "teacher@test.com",
		HashedPassword: auth.HashPassword("12345678"),
		Role:           core.Teacher,
	}
	var userStudent = core.User{
		FirstName:      "student",
		LastName:       "student",
		Email:          "student@test.com",
		HashedPassword: auth.HashPassword("12345678"),
		Role:           core.Student,
	}
	BeforeAll(func() {
		config.LoadConfigTest()
		config.SetupInstance()
		handlers.SetupRouter()
		db.UserDBInstance.Clear()
		db.JwtDBInstance.Clear()
		db.CourseDBInstance.Clear()
		db.LessonDBInstance.Clear()
		db.NodeDBInstance.Clear()
		mockUser := []core.User{userAdmin, userTeacher, userStudent}
		err := db.UserDBInstance.InsertManyUser(mockUser)
		userAdmin = mockUser[0]
		userTeacher = mockUser[1]
		userStudent = mockUser[2]
		Expect(err).Should(BeNil())
		t = GinkgoT()
	})

	var token = core.Token{}
	Context("get user", func() {
		It("get token", func() {
			res := apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/sign-in").
				JSON(`
				{
					"email": "admin@test.com",
					"password": "12345678"
				}
				`).
				Expect(t).
				CookiePresent("refresh_token").
				Status(http.StatusOK).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("firstName", "admin").
						Equal("lastName", "admin").
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

		lessonID := ""
		It("post lesson", func() {
			req := apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/auth/lessons").
				Header("Authorization", "Bearer "+token.Access).
				JSON(`
				{
					"name": "course name",
					"description": "course description",
					"access":"public",
					"lessons": [],
					"nodes": [],
					"nextLesson": [],
					"prevLesson": []
				}`).
				Expect(t).
				Status(http.StatusCreated).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("name", "course name").
						Equal("description", "course description").
						End(),
				).
				End().Response
			body := getJSON(req)
			lessonID = body["id"].(string)
		})

		It("get lessons", func() {
			apitest.New().
				Handler(handlers.Router).
				Get("/api/v1/auth/lessons/"+lessonID).
				Header("Authorization", "Bearer "+token.Access).
				Expect(t).
				Status(http.StatusOK).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("name", "course name").
						Equal("description", "course description").
						End(),
				).
				End()
		})

		It("update lessons", func() {
			apitest.New().
				Handler(handlers.Router).
				Patch("/api/v1/auth/lessons/"+lessonID).
				Header("Authorization", "Bearer "+token.Access).
				JSON(`
				{
					"name": "course name update",
					"description": "course description update"
				}
				`).
				Expect(t).
				Status(http.StatusOK).
				Assert(
					jsonpath.Chain().
						Present("id").
						Equal("name", "course name update").
						Equal("description", "course description update").
						End(),
				).
				End()
		})

		It("delete lessons", func() {
			apitest.New().
				Handler(handlers.Router).
				Delete("/api/v1/auth/lessons/"+lessonID).
				Header("Authorization", "Bearer "+token.Access).
				Expect(t).
				Status(http.StatusOK).
				End()
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
