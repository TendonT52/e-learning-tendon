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

var _ = Describe("test lesson db", Ordered, func() {
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

	Context("test create course", func() {
		When("success", func() {
			var user core.User
			var nodes []core.Node
			BeforeEach(func() {
				user = core.User{
					FirstName:      "testFirstName",
					LastName:       "testLastName",
					Email:          "testEmail",
					HashedPassword: "testHashPassword",
					Role:           core.Student,
					Courses:        []string{},
				}
				err := db.UserDBInstance.InsertUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				nodes = []core.Node{
					{
						Type:     core.Text,
						Data:     "hello",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Image,
						Data:     "path/to/image",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Pdf,
						Data:     "path/to/pdf",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
				}
				err = db.NodeDBInstance.InsertManyNode(nodes)
				Expect(err).Should(BeNil())
			})
			It("should success", func() {
				lesson := core.Lesson{
					Name:        "testLesson",
					Description: "testDescription",
					Access:      core.PublicAccess,
					Nodes:       []string{nodes[0].ID, nodes[1].ID, nodes[2].ID},
					NextLessons: []string{},
					PrevLessons: []string{},
					CreateBy:    user.ID,
				}
				err := app.CreateLesson(&lesson)
				Expect(err).Should(BeNil())
				Expect(lesson).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Not(BeEmpty()),
					"Name":        Equal("testLesson"),
					"Description": Equal("testDescription"),
					"Access":      Equal(core.PublicAccess),
					"Nodes":       Equal([]string{nodes[0].ID, nodes[1].ID, nodes[2].ID}),
					"NextLessons": Equal([]string{}),
					"PrevLessons": Equal([]string{}),
					"CreateBy":    Equal(user.ID),
					"UpdatedAt":   Not(BeNil()),
				}))
			})
		})
	})

	Context("test get lesson", func() {
		When("success", func() {
			var user core.User
			var nodes []core.Node
			var lesson core.Lesson
			BeforeEach(func() {
				user = core.User{
					FirstName:      "testFirstName",
					LastName:       "testLastName",
					Email:          "testEmail",
					HashedPassword: "testHashPassword",
					Role:           core.Student,
					Courses:        []string{},
				}
				err := db.UserDBInstance.InsertUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				nodes = []core.Node{
					{
						Type:     core.Text,
						Data:     "hello",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Image,
						Data:     "path/to/image",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Pdf,
						Data:     "path/to/pdf",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
				}
				err = db.NodeDBInstance.InsertManyNode(nodes)
				Expect(err).Should(BeNil())
				lesson = core.Lesson{
					Name:        "testLesson",
					Description: "testDescription",
					Access:      core.PublicAccess,
					Nodes:       []string{nodes[0].ID, nodes[1].ID, nodes[2].ID},
					NextLessons: []string{},
					PrevLessons: []string{},
					CreateBy:    user.ID,
				}
				err = app.CreateLesson(&lesson)
				Expect(err).Should(BeNil())
			})
			It("should success", func() {
				lesson, err := app.GetLesson(lesson.ID)
				Expect(err).Should(BeNil())
				Expect(lesson).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(lesson.ID),
					"Name":        Equal("testLesson"),
					"Description": Equal("testDescription"),
					"Access":      Equal(core.PublicAccess),
					"Nodes":       Equal([]string{nodes[0].ID, nodes[1].ID, nodes[2].ID}),
					"NextLessons": Equal([]string{}),
					"PrevLessons": Equal([]string{}),
					"CreateBy":    Equal(user.ID),
					"UpdatedAt":   Not(BeNil()),
				}))
			})
		})
	})

	Context("test update lesson", func() {
		When("success", func() {
			var user core.User
			var nodes []core.Node
			var lesson core.Lesson
			BeforeEach(func() {
				user = core.User{
					FirstName:      "testFirstName",
					LastName:       "testLastName",
					Email:          "testEmail",
					HashedPassword: "testHashPassword",
					Role:           core.Student,
					Courses:        []string{},
				}
				err := db.UserDBInstance.InsertUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				nodes = []core.Node{
					{
						Type:     core.Text,
						Data:     "hello",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Image,
						Data:     "path/to/image",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Pdf,
						Data:     "path/to/pdf",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
				}
				err = db.NodeDBInstance.InsertManyNode(nodes)
				Expect(err).Should(BeNil())
				lesson = core.Lesson{
					Name:        "testLesson",
					Description: "testDescription",
					Access:      core.PublicAccess,
					Nodes:       []string{nodes[0].ID, nodes[1].ID, nodes[2].ID},
					NextLessons: []string{},
					PrevLessons: []string{},
					CreateBy:    user.ID,
				}
				err = app.CreateLesson(&lesson)
				Expect(err).Should(BeNil())
			})
			It("should success", func() {
				lesson.Name = "testLesson2"
				lesson.Description = "testDescription2"
				lesson.Access = core.PrivateAccess
				lesson.Nodes = []string{nodes[0].ID, nodes[1].ID}
				lesson.NextLessons = []string{lesson.ID}
				lesson.PrevLessons = []string{lesson.ID}
				err := app.UpdateLesson(&lesson)
				Expect(err).Should(BeNil())
				lesson, err := app.GetLesson(lesson.ID)
				Expect(err).Should(BeNil())
				Expect(lesson).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(lesson.ID),
					"Name":        Equal("testLesson2"),
					"Description": Equal("testDescription2"),
					"Access":      Equal(core.PrivateAccess),
					"Nodes":       Equal([]string{nodes[0].ID, nodes[1].ID}),
					"NextLessons": Equal([]string{lesson.ID}),
					"PrevLessons": Equal([]string{lesson.ID}),
					"CreateBy":    Equal(user.ID),
					"UpdatedAt":   Not(BeNil()),
				}))
			})
		})
	})

	Context("test delete lesson", func() {
		When("success", func() {
			var user core.User
			var nodes []core.Node
			var lesson core.Lesson
			BeforeEach(func() {
				user = core.User{
					FirstName:      "testFirstName",
					LastName:       "testLastName",
					Email:          "testEmail",
					HashedPassword: "testHashPassword",
					Role:           core.Student,
					Courses:        []string{},
				}
				err := db.UserDBInstance.InsertUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				nodes = []core.Node{
					{
						Type:     core.Text,
						Data:     "hello",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Image,
						Data:     "path/to/image",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
					{
						Type:     core.Pdf,
						Data:     "path/to/pdf",
						CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
					},
				}
				err = db.NodeDBInstance.InsertManyNode(nodes)
				Expect(err).Should(BeNil())
				lesson = core.Lesson{
					Name:        "testLesson",
					Description: "testDescription",
					Access:      core.PublicAccess,
					Nodes:       []string{nodes[0].ID, nodes[1].ID, nodes[2].ID},
					NextLessons: []string{},
					PrevLessons: []string{},
					CreateBy:    user.ID,
				}
				err = app.CreateLesson(&lesson)
				Expect(err).Should(BeNil())
			})
			It("should success", func() {
				err := app.DeleteLesson(lesson.ID)
				Expect(err).Should(BeNil())
				lesson, err := app.GetLesson(lesson.ID)
				Expect(err).Should(MatchError(errs.NewHttpError(http.StatusNotFound, "lesson id is not found")))
				Expect(lesson).Should(BeZero())
			})
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
