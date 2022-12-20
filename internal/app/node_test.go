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

var _ = Describe("test node db", Ordered, func() {
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

	Context("test create node", func() {
		var user core.User
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
		})
		It("should create node successfully", func() {
			node := core.Node{
				Type:     core.Text,
				Data:     "testData",
				CreateBy: user.ID,
			}
			err := app.CreateNode(&node)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(node).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Not(BeEmpty()),
				"Type":      Equal(core.Text),
				"Data":      Equal("testData"),
				"CreateBy":  Equal(user.ID),
				"UpdatedAt": Not(BeZero()),
			}))
		})
	})

	Context("test get node", func() {
		var user core.User
		var node core.Node
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
			node = core.Node{
				Type:     core.Text,
				Data:     "testData",
				CreateBy: user.ID,
			}
			err = app.CreateNode(&node)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should get node successfully", func() {
			node, err := app.GetNode(node.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(node).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(node.ID),
				"Type":      Equal(core.Text),
				"Data":      Equal("testData"),
				"CreateBy":  Equal(user.ID),
				"UpdatedAt": Not(BeZero()),
			}))
		})
	})

	Context("test update node", func() {
		var user core.User
		var node core.Node
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
			node = core.Node{
				Type:     core.Text,
				Data:     "testData",
				CreateBy: user.ID,
			}
			err = app.CreateNode(&node)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should update node successfully", func() {
			node.Data = "testDataUpdate"
			err := app.UpdateNode(&node)
			Expect(err).ShouldNot(HaveOccurred())
			node, err = app.GetNode(node.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(node).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(node.ID),
				"Type":      Equal(core.Text),
				"Data":      Equal("testDataUpdate"),
				"CreateBy":  Equal(user.ID),
				"UpdatedAt": Not(BeZero()),
			}))
		})
	})

	Context("test delete node", func() {
		var user core.User
		var node core.Node
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
			node = core.Node{
				Type:     core.Text,
				Data:     "testData",
				CreateBy: user.ID,
			}
			err = app.CreateNode(&node)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should delete node successfully", func() {
			err := app.DeleteNode(node.ID)
			Expect(err).ShouldNot(HaveOccurred())
			_, err = app.GetNode(node.ID)
			Expect(err).Should(MatchError(errs.NewHttpError(http.StatusNotFound, "node id is not found")))
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
