package app_test

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/app"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("test course db", Ordered, func() {
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
			var lessons []core.Lesson
			BeforeEach(func() {
				user = core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				lessons = []core.Lesson{
					{
						Name:        "Lesson 1",
						Description: "description 1",
						Access:      core.PublicAccess,
						CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
						Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
					},
					{
						Name:        "Lesson 2",
						Description: "description 2",
						Access:      core.ProtectedAccess,
						CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
						Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
					},
					{
						Name:        "Lesson 3",
						Description: "description 3",
						Access:      core.PrivateAccess,
						CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
						Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
						PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
					},
				}
				err = db.LessonDBInstance.InsertManyLesson(lessons)
				Expect(err).To(BeNil())
				Expect(lessons).To(HaveLen(3))
			})
			It("should create course", func() {
				course := core.Course{
					Name:        "testcourse",
					Description: "testcourse",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{lessons[0].ID, lessons[1].ID, lessons[2].ID},
				}
				err := app.CreateCourse(&course)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(course.ID).ShouldNot(BeEmpty())
				courseCheck, err := db.CourseDBInstance.FindCourse(course.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(courseCheck).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(course.ID),
					"Name":        Equal(course.Name),
					"Description": Equal(course.Description),
					"Access":      Equal(course.Access),
					"CreateBy":    Equal(course.CreateBy),
					"UpdatedAt":   Equal(course.UpdatedAt),
					"Lessons":     Equal([]string{lessons[0].ID, lessons[1].ID, lessons[2].ID}),
				}))
			})
		})
	})

	Context("test get course", func() {
		var user core.User
		var course core.Course
		When("success", func() {
			BeforeEach(func() {
				user = core.User{
					FirstName: "firstnametest1",
					LastName:  "lastnametest1",
					Email:     "email@test.com",
					Role:      core.Student,
					Courses:   []string{},
				}
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				course = core.Course{
					Name:        "testcourse",
					Description: "testcourse",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{},
				}
				err = app.CreateCourse(&course)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(course.ID).ShouldNot(BeZero())
			})
			It("should get course", func() {
				courseCheck, err := app.GetCourse(course.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(courseCheck).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(course.ID),
					"Name":        Equal(course.Name),
					"Description": Equal(course.Description),
					"Access":      Equal(course.Access),
					"CreateBy":    Equal(course.CreateBy),
					"UpdatedAt":   Equal(course.UpdatedAt),
					"Lessons":     Equal([]string{}),
				}))
			})

		})
	})

	Context("test update course", func() {
		When("success", func() {
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
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				course = core.Course{
					Name:        "testcourse",
					Description: "testcourse",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{},
				}
				err = app.CreateCourse(&course)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(course.ID).ShouldNot(BeZero())
			})
			It("should update course", func() {
				course.Name = "testcourseupdate"
				course.Description = "testcourseupdate"
				course.Access = core.PrivateAccess
				err := app.UpdateCourse(&course)
				Expect(err).ShouldNot(HaveOccurred())
				courseCheck, err := db.CourseDBInstance.FindCourse(course.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(courseCheck).Should(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(course.ID),
					"Name":        Equal(course.Name),
					"Description": Equal(course.Description),
					"Access":      Equal(course.Access),
					"CreateBy":    Equal(course.CreateBy),
					"UpdatedAt":   Not(BeNil()),
					"Lessons":     Equal([]string{}),
				}))
			})
		})
	})

	Context("test delete course", func() {
		When("success", func() {
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
				token, err := app.SignUp(&user, "password1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(token).ShouldNot(BeNil())
				course = core.Course{
					Name:        "testcourse",
					Description: "testcourse",
					Access:      core.PublicAccess,
					CreateBy:    user.ID,
					Lessons:     []string{},
				}
				err = app.CreateCourse(&course)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(course.ID).ShouldNot(BeZero())
			})
			It("should delete course", func() {
				err := app.DeleteCourse(course.ID)
				Expect(err).ShouldNot(HaveOccurred())
				courseCheck, err := db.CourseDBInstance.FindCourse(course.ID)
				Expect(err).Should(HaveOccurred())
				Expect(courseCheck).Should(BeZero())
			})
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
