package app_test

import (
	"net/http"
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("test user db", func() {
	BeforeEach(func() {
		db.NewClient("mongodb://admin:password@localhost:27017",
			db.MongoConfig{
				InsertTimeOut: time.Minute,
				FindTimeOut:   time.Minute,
				UpdateTimeOut: time.Minute,
				DeleteTimeOut: time.Minute,
			})
		db.NewDB("tendon")
		db.NewUserDB("user_test")
		db.NewJwtDB("jwt_test")
		db.NewCourseDB("course_test")
		db.NewLessonDB("lesson_test")
		db.NewNodeDB("node_test")
		db.UserDBInstance.Clear()
		db.JwtDBInstance.Clear()
		db.CourseDBInstance.Clear()
		db.LessonDBInstance.Clear()
		db.NodeDBInstance.Clear()
		appConfig := app.AppConfig{
			AppName:              "issuerTest",
			AccessSecret:         "jwtAccessSecret",
			RefreshSecret:        "jwtRefreshSecret",
			AccesstokenDuration:  time.Minute,
			RefreshtokenDuration: time.Minute * 5,
		}

		reposConfig := app.ReposInstance{
			UserDB:   db.UserDBInstance,
			JwtDB:    db.JwtDBInstance,
			CourseDB: db.CourseDBInstance,
			LessonDB: db.LessonDBInstance,
			NodeDB:   db.NodeDBInstance,
		}

		app.NewApp(appConfig, reposConfig)
	})

	Context("test sign up", func() {
		When("sign up with valid email and password", func() {
			It("success", func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Not(BeEmpty()),
					"FirstName":      Equal("firstnametest1"),
					"LastName":       Equal("lastnametest1"),
					"Email":          Equal("email@test.com"),
					"HashedPassword": Not(BeEmpty()),
					"Role":           Equal(core.Student),
					"Courses":        Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
		When("sign up with already use email", func() {
			BeforeEach(func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
			})
			It("fail", func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(token).Should(BeZero())
				Expect(err).Should(MatchError(errs.NewHttpError(http.StatusConflict, "email already exists")))
			})

		})
	})

	Context("test sign in", func() {
		When("sign in with valid email and password", func() {
			BeforeEach(func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
			})

			It("success", func() {
				user, token, err := app.SignIn("email@test.com", "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Not(BeEmpty()),
					"FirstName":      Equal("firstnametest1"),
					"LastName":       Equal("lastnametest1"),
					"Email":          Equal("email@test.com"),
					"HashedPassword": Not(BeEmpty()),
					"Role":           Equal(core.Student),
					"Courses":        Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
		When("sign in with invalid email or password", func() {
			BeforeEach(func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
			})
			It("fail invalid email", func() {
				user, token, err := app.SignIn("invalidEmail", "password1")
				Expect(user).Should(BeZero())
				Expect(token).Should(BeZero())
				Expect(err).Should(MatchError(errs.NewHttpError(http.StatusUnauthorized, "invalid email or password")))
			})
		})
	})

	Context("sign out test", func() {
		When("sign out with valid token", func() {
			var token core.Token
			BeforeEach(func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				var err error
				token, err = app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
			})
			It("success", func() {
				err := app.SignOut(token)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Context("test get user", func() {
		When("valid get user", func() {
			var userID string
			BeforeEach(func() {
				user := core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				var err error
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				userID = user.ID
			})
			It("success", func() {
				user, err := app.GetUser(userID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userID),
					"FirstName":      Equal("firstnametest1"),
					"LastName":       Equal("lastnametest1"),
					"Email":          Equal("email@test.com"),
					"HashedPassword": Not(BeEmpty()),
					"Role":           Equal(core.Student),
					"Courses":        Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
	})

	Context("test update user", func() {
		When("valid update user", func() {
			var user core.User
			var course core.Course
			BeforeEach(func() {
				user = core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				var err error
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				course = core.Course{
					Name:        "testcourse1",
					Description: "testcourse1",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{},
				}
				db.CourseDBInstance.InsertCourse(&course)
			})
			It("success", func() {
				userUpdate := core.User{
					ID:        user.ID,
					FirstName: "updatefirstnametest1",
					LastName:  "updatetestlastnametest1",
					Email:     "updateemail@test.com",
					Role:      core.Student,
					Courses:   []string{course.ID},
				}
				err := app.UpdateUser(&userUpdate, "")
				Expect(err).ShouldNot(HaveOccurred())
				userCheck, err := db.UserDBInstance.FindUser(user.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(userCheck).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(user.ID),
					"FirstName":      Equal("updatefirstnametest1"),
					"LastName":       Equal("updatetestlastnametest1"),
					"Email":          Equal("updateemail@test.com"),
					"HashedPassword": Not(BeEmpty()),
					"Role":           Equal(core.Student),
					"Courses":        Equal([]string{course.ID}),
					"UpdatedAt":      Not(BeNil()),
				}))

			})
		})
	})

	Context("test delete user", func() {
		When("valid delete user", func() {
			var user core.User
			var course core.Course
			BeforeEach(func() {
				user = core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				var err error
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				course = core.Course{
					Name:        "testcourse1",
					Description: "testcourse1",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{},
				}
				db.CourseDBInstance.InsertCourse(&course)
			})
			It("success", func() {
				err := app.DeleteUser(user.ID)
				Expect(err).ShouldNot(HaveOccurred())
				_, err = db.UserDBInstance.FindUser(user.ID)
				Expect(err).Should(MatchError(errs.UserNotFound))
			})
		})
	})

	AfterEach(func() {
		db.DisconnectMongo()
	})
})
