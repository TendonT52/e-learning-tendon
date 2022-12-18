package db_test

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Insert node to db", Ordered, func() {
	BeforeEach(func() {
		db.NewClient("mongodb://admin:password@localhost:27017",
			db.MongoConfig{
				InsertTimeOut: time.Minute,
				FindTimeOut:   time.Minute,
				UpdateTimeOut: time.Minute,
				DeleteTimeOut: time.Minute,
			})
		db.NewDB("tendon")
		db.NewLessonDB("node_test")
		db.LessonDBInstance.Clear()
	})

	Context("Insert one node", func() {
		It("Should insert one node", func() {
			lesson := core.Lesson{
				Name:        "Lesson 1",
				Description: "Lesson 1",
				Access:      "public",
				CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
				Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
			}
			err := db.LessonDBInstance.InsertLesson(&lesson)
			Expect(err).To(BeNil())
			Expect(lesson).To(MatchFields(IgnoreExtras, Fields{
				"ID":          Not(BeEmpty()),
				"Name":        Equal("Lesson 1"),
				"Description": Equal("Lesson 1"),
				"Access":      Equal("public"),
				"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
				"UpdatedAt":   Not(BeNil()),
				"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
			}))
		})
	})

	Context("Insert many nodes", func() {
		It("Should insert multiple nodes", func() {
			lessons := []core.Lesson{
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
			err := db.LessonDBInstance.InsertManyLesson(lessons)
			Expect(err).To(BeNil())
			Expect(lessons).To(HaveLen(3))
			Expect(lessons[0]).To(MatchFields(IgnoreExtras, Fields{
				"ID":          Not(BeEmpty()),
				"Name":        Equal("Lesson 1"),
				"Description": Equal("description 1"),
				"Access":      Equal(core.PublicAccess),
				"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
				"UpdatedAt":   Not(BeNil()),
				"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
			}))
			Expect(lessons[1]).To(MatchFields(IgnoreExtras, Fields{
				"ID":          Not(BeEmpty()),
				"Name":        Equal("Lesson 2"),
				"Description": Equal("description 2"),
				"Access":      Equal(core.ProtectedAccess),
				"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
				"UpdatedAt":   Not(BeNil()),
				"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
			}))
			Expect(lessons[2]).To(MatchFields(IgnoreExtras, Fields{
				"ID":          Not(BeEmpty()),
				"Name":        Equal("Lesson 3"),
				"Description": Equal("description 3"),
				"Access":      Equal(core.PrivateAccess),
				"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
				"UpdatedAt":   Not(BeNil()),
				"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
			}))
		})
	})

	Context("Find lesson", func() {
		lessonIDs := make([]string, 3)
		BeforeEach(func() {
			lessons := []core.Lesson{
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
			err := db.LessonDBInstance.InsertManyLesson(lessons)
			Expect(err).To(BeNil())
			lessonIDs[0] = lessons[0].ID
			lessonIDs[1] = lessons[1].ID
			lessonIDs[2] = lessons[2].ID
		})
		When("success", func() {
			It("should return lesson", func() {
				lesson, err := db.LessonDBInstance.FindLesson(lessonIDs[1])
				Expect(err).To(BeNil())
				Expect(lesson).To(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(lessonIDs[1]),
					"Name":        Equal("Lesson 2"),
					"Description": Equal("description 2"),
					"Access":      Equal(core.ProtectedAccess),
					"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
					"UpdatedAt":   Not(BeNil()),
					"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				}))
			})
		})
	})

	Context("find many lesson", func() {
		lessonIDs := make([]string, 3)
		BeforeEach(func() {
			lessons := []core.Lesson{
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
			err := db.LessonDBInstance.InsertManyLesson(lessons)
			Expect(err).To(BeNil())
			lessonIDs[0] = lessons[0].ID
			lessonIDs[1] = lessons[1].ID
			lessonIDs[2] = lessons[2].ID
		})
		When("success", func() {
			It("should return lesson", func() {
				lessons, err := db.LessonDBInstance.FindManyLesson(lessonIDs)
				Expect(err).To(BeNil())
				Expect(lessons).To(HaveLen(3))
				Expect(lessons).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"ID":          Equal(lessonIDs[0]),
						"Name":        Equal("Lesson 1"),
						"Description": Equal("description 1"),
						"Access":      Equal(core.PublicAccess),
						"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
						"UpdatedAt":   Not(BeNil()),
						"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					}),
					MatchFields(IgnoreExtras, Fields{
						"ID":          Equal(lessonIDs[1]),
						"Name":        Equal("Lesson 2"),
						"Description": Equal("description 2"),
						"Access":      Equal(core.ProtectedAccess),
						"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
						"UpdatedAt":   Not(BeNil()),
						"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					}),
					MatchFields(IgnoreExtras, Fields{
						"ID":          Equal(lessonIDs[2]),
						"Name":        Equal("Lesson 3"),
						"Description": Equal("description 3"),
						"Access":      Equal(core.PrivateAccess),
						"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
						"UpdatedAt":   Not(BeNil()),
						"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
						"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					}),
				))
			})
		})
	})

	Context("update lesson", func() {
		var lessonID string
		BeforeEach(func() {
			lesson := &core.Lesson{
				Name:        "Lesson 1",
				Description: "description 1",
				Access:      core.PublicAccess,
				CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
				Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
			}
			err := db.LessonDBInstance.InsertLesson(lesson)
			Expect(err).To(BeNil())
			lessonID = lesson.ID
		})
		When("success", func() {
			It("should return lesson", func() {
				lesson := core.Lesson{
					ID:          lessonID,
					Name:        "update lesson 2",
					Description: "update description 2",
					Access:      core.ProtectedAccess,
					CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
					Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
					NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
					PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				}
				err := db.LessonDBInstance.UpdateLesson(&lesson)
				Expect(err).To(BeNil())
				Expect(lesson).To(MatchFields(IgnoreExtras, Fields{
					"ID":          Equal(lessonID),
					"Name":        Equal("update lesson 2"),
					"Description": Equal("update description 2"),
					"Access":      Equal(core.ProtectedAccess),
					"CreateBy":    Equal("5f7f7d1e1b8a7c2c2c0a7a1d"),
					"UpdatedAt":   Not(BeNil()),
					"Nodes":       Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					"NextLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
					"PrevLessons": Equal([]string{"5f7f7d1e1b8a7c2c2c0a7a1d"}),
				}))
			})
		})
	})

	Context("delete lesson", func() {
		var lessonID string
		BeforeEach(func() {
			lesson := &core.Lesson{
				Name:        "Lesson 1",
				Description: "description 1",
				Access:      core.PublicAccess,
				CreateBy:    "5f7f7d1e1b8a7c2c2c0a7a1d",
				Nodes:       []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				NextLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
				PrevLessons: []string{"5f7f7d1e1b8a7c2c2c0a7a1d"},
			}
			err := db.LessonDBInstance.InsertLesson(lesson)
			Expect(err).To(BeNil())
			lessonID = lesson.ID
		})
		When("success", func() {
			It("should return nil", func() {
				err := db.LessonDBInstance.DeleteLesson(lessonID)
				Expect(err).To(BeNil())
				lesson, err := db.LessonDBInstance.FindLesson(lessonID)
				Expect(err).To(MatchError(errs.LessonNotFound))
				Expect(lesson).To(BeZero())
			})
		})
	})

	Context("delete lesson", func() {
		lessonIDs := make([]string, 3)
		BeforeEach(func() {
			lessons := []core.Lesson{
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
			err := db.LessonDBInstance.InsertManyLesson(lessons)
			Expect(err).To(BeNil())
			lessonIDs[0] = lessons[0].ID
			lessonIDs[1] = lessons[1].ID
			lessonIDs[2] = lessons[2].ID
		})
		When("success", func() {
			It("should return nil", func() {
				err := db.LessonDBInstance.DeleteManyLesson(lessonIDs)
				Expect(err).To(BeNil())
				lessons, err := db.LessonDBInstance.FindManyLesson(lessonIDs)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(lessons).To(HaveLen(0))
			})
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
