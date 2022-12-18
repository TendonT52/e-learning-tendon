package db_test

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/core"
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
		db.NewCurriculumDB("curricula_test")
		db.CurriculumDBInstance.Clear()
	})
	Context("Insert one curriculum", func() {
		It("Should insert one curriculum", func() {
			curriculum := core.Curriculum{
				Name:        "test",
				Description: "test",
				Access:      core.PublicAccess,
				CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
				Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
			}
			err := db.CurriculumDBInstance.InsertCurriculum(&curriculum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curriculum).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Not(BeEmpty()),
				"UpdatedAt": Not(BeNil()),
			}))
		})
	})

	Context("Insert many curriculum", func() {
		It("Should insert many curriculum", func() {
			curriculums := []core.Curriculum{
				{
					Name:        "test",
					Description: "test",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test",
					Description: "test",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curriculums)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curriculums).Should(ConsistOf(
				MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeEmpty()),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeEmpty()),
					"UpdatedAt": Not(BeNil()),
				}),
			))
		})
	})

	Context("find curriculum", func() {
		curriculumIDs := make([]string, 3)
		BeforeEach(func() {
			curricula := []core.Curriculum{
				{
					Name:        "test1",
					Description: "test1",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test2",
					Description: "test2",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test3",
					Description: "test3",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curricula)
			Expect(err).ShouldNot(HaveOccurred())
			for i := range curricula {
				curriculumIDs[i] = curricula[i].ID
			}
		})
		It("Should find curriculum by id", func() {
			curriculum, err := db.CurriculumDBInstance.FindCurriculum(curriculumIDs[0])
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curriculum).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(curriculumIDs[0]),
				"Name":      Equal("test1"),
				"Access":    Equal(core.PublicAccess),
				"CreateBy":  Equal("5f9f1b5b5d1c3b0b8c1c1c1c"),
				"Lessons":   Equal([]string{"5f9f1b5b5d1c3b0b8c1c1c1c"}),
				"UpdatedAt": Not(BeNil()),
			}))
		})
	})

	Context("find many curricula", func() {
		curriculumIDs := make([]string, 3)
		BeforeEach(func() {
			curricula := []core.Curriculum{
				{
					Name:        "test1",
					Description: "test1",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test2",
					Description: "test2",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test3",
					Description: "test3",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curricula)
			Expect(err).ShouldNot(HaveOccurred())
			for i := range curricula {
				curriculumIDs[i] = curricula[i].ID
			}
		})
		It("Should find many curricula by ids", func() {
			curricula, err := db.CurriculumDBInstance.FindManyCurriculum(curriculumIDs)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curricula).Should(ConsistOf(
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(curriculumIDs[0]),
					"Name":      Equal("test1"),
					"Access":    Equal(core.PublicAccess),
					"CreateBy":  Equal("5f9f1b5b5d1c3b0b8c1c1c1c"),
					"Lessons":   Equal([]string{"5f9f1b5b5d1c3b0b8c1c1c1c"}),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(curriculumIDs[1]),
					"Name":      Equal("test2"),
					"Access":    Equal(core.PublicAccess),
					"CreateBy":  Equal("5f9f1b5b5d1c3b0b8c1c1c1c"),
					"Lessons":   Equal([]string{"5f9f1b5b5d1c3b0b8c1c1c1c"}),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(curriculumIDs[2]),
					"Name":      Equal("test3"),
					"Access":    Equal(core.PublicAccess),
					"CreateBy":  Equal("5f9f1b5b5d1c3b0b8c1c1c1c"),
					"Lessons":   Equal([]string{"5f9f1b5b5d1c3b0b8c1c1c1c"}),
					"UpdatedAt": Not(BeNil()),
				}),
			))
		})
	})

	Context("update curriculum", func() {
		curriculumIDs := make([]string, 3)
		BeforeEach(func() {
			curricula := []core.Curriculum{
				{
					Name:        "test1",
					Description: "test1",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test2",
					Description: "test2",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test3",
					Description: "test3",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curricula)
			Expect(err).ShouldNot(HaveOccurred())
			for i := range curricula {
				curriculumIDs[i] = curricula[i].ID
			}
		})
		It("Should update curriculum", func() {
			curriculum := core.Curriculum{
				ID:          curriculumIDs[0],
				Name:        "update test1",
				Description: "update test1",
				Access:      core.PublicAccess,
				CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
				Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
			}
			err := db.CurriculumDBInstance.UpdateCurriculum(&curriculum)
			Expect(err).ShouldNot(HaveOccurred())
			curriculum, err = db.CurriculumDBInstance.FindCurriculum(curriculumIDs[0])
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curriculum).Should(MatchFields(IgnoreExtras, Fields{
				"ID":           Equal(curriculumIDs[0]),
				"Name":         Equal("update test1"),
				"Description": Equal("update test1"),
				"Access":       Equal(core.PublicAccess),
				"CreateBy":     Equal("5f9f1b5b5d1c3b0b8c1c1c1c"),
				"Lessons":      Equal([]string{"5f9f1b5b5d1c3b0b8c1c1c1c"}),
				"UpdatedAt":    Not(BeNil()),
			}))
		})
	})

	Context("delete curriculum", func() {
		curriculumIDs := make([]string, 3)
		BeforeEach(func() {
			curricula := []core.Curriculum{
				{
					Name:        "test1",
					Description: "test1",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test2",
					Description: "test2",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test3",
					Description: "test3",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curricula)
			Expect(err).ShouldNot(HaveOccurred())
			for i := range curricula {
				curriculumIDs[i] = curricula[i].ID
			}
		})
		It("Should delete curriculum", func() {
			err := db.CurriculumDBInstance.DeleteCurriculum(curriculumIDs[0])
			Expect(err).ShouldNot(HaveOccurred())
			curriculum, err := db.CurriculumDBInstance.FindCurriculum(curriculumIDs[0])
			Expect(err).Should(HaveOccurred())
			Expect(curriculum).Should(BeZero())
		})
	})

	Context("delete may curricula", func() {
		curriculumIDs := make([]string, 3)
		BeforeEach(func() {
			curricula := []core.Curriculum{
				{
					Name:        "test1",
					Description: "test1",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test2",
					Description: "test2",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
				{
					Name:        "test3",
					Description: "test3",
					Access:      core.PublicAccess,
					CreateBy:    "5f9f1b5b5d1c3b0b8c1c1c1c",
					Lessons:     []string{"5f9f1b5b5d1c3b0b8c1c1c1c"},
				},
			}
			err := db.CurriculumDBInstance.InsertManyCurriculum(curricula)
			Expect(err).ShouldNot(HaveOccurred())
			for i := range curricula {
				curriculumIDs[i] = curricula[i].ID
			}
		})
		It("Should delete may curricula", func() {
			err := db.CurriculumDBInstance.DeleteManyCurriculum(curriculumIDs)
			Expect(err).ShouldNot(HaveOccurred())
			curricula, err := db.CurriculumDBInstance.FindManyCurriculum(curriculumIDs)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(curricula).Should(HaveLen(0))
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})

