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

var _ = Describe("test user db", Ordered, func() {
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
		db.UserDBInstance.Clear()
	})

	Context("Insert user to db", func() {
		When("Success", func() {
			It("should return user", func() {
				user := core.User{
					FirstName:      "testFirstName",
					LastName:       "testLastName",
					Email:          "testEmail",
					HashedPassword: "testHashPassword",
					Role:           core.Student,
					Courses:        []string{},
				}
				err := db.UserDBInstance.InsertUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Not(BeNil()),
					"FirstName":      Equal("testFirstName"),
					"LastName":       Equal("testLastName"),
					"Email":          Equal("testEmail"),
					"HashedPassword": Equal("testHashPassword"),
					"Role":           Equal(core.Student),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
	})

	Context("Insert many user to db", func() {
		When("Success", func() {
			It("should return user", func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"ID":             Not(BeZero()),
						"FirstName":      Equal("testFirstName1"),
						"LastName":       Equal("testLastName1"),
						"Email":          Equal("testEmail1"),
						"HashedPassword": Equal("testHashPassword1"),
						"Role":           Equal(core.Student),
						"Courses":      Equal([]string{}),
						"UpdatedAt":      Not(BeNil()),
					}),
					MatchFields(IgnoreExtras, Fields{
						"ID":             Not(BeZero()),
						"FirstName":      Equal("testFirstName2"),
						"LastName":       Equal("testLastName2"),
						"Email":          Equal("testEmail2"),
						"HashedPassword": Equal("testHashPassword2"),
						"Role":           Equal(core.Teacher),
						"Courses":      Equal([]string{}),
						"UpdatedAt":      Not(BeNil()),
					}),
					MatchFields(IgnoreExtras, Fields{
						"ID":             Not(BeZero()),
						"FirstName":      Equal("testFirstName3"),
						"LastName":       Equal("testLastName3"),
						"Email":          Equal("testEmail3"),
						"HashedPassword": Equal("testHashPassword3"),
						"Role":           Equal(core.Admin),
						"Courses":      Equal([]string{}),
						"UpdatedAt":      Not(BeNil()),
					}),
				))
			})
		})
	})

	Context("Find user by email", func() {
		When("Success", func() {
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
			})
			It("should return user", func() {
				user, err := db.UserDBInstance.FindUserByEmail("testEmail2")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Not(BeNil()),
					"FirstName":      Equal("testFirstName2"),
					"LastName":       Equal("testLastName2"),
					"Email":          Equal("testEmail2"),
					"HashedPassword": Equal("testHashPassword2"),
					"Role":           Equal(core.Teacher),
					"Courses":      Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
		When("user not found", func() {
			It("should return nil", func() {
				user, err := db.UserDBInstance.FindUserByEmail("testEmail2")
				Expect(err).To(MatchError(errs.UserNotFound))
				Expect(user).To(BeZero())
			})
		})
	})

	Context("Find user", func() {
		userIDs := make([]string, 3)
		When("Success", func() {
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
				userIDs[2] = users[2].ID
			})
			It("should return user", func() {
				user, err := db.UserDBInstance.FindUser(userIDs[1])
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userIDs[1]),
					"FirstName":      Equal("testFirstName2"),
					"LastName":       Equal("testLastName2"),
					"Email":          Equal("testEmail2"),
					"HashedPassword": Equal("testHashPassword2"),
					"Role":           Equal(core.Teacher),
					"Courses":      Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
		When("Fail", func() {
			It("should return user not found", func() {
				user, err := db.UserDBInstance.FindUser(userIDs[1])
				Expect(err).To(MatchError(errs.UserNotFound))
				Expect(user).To(BeZero())
			})
			It("should return invalid id", func() {
				_, err := db.UserDBInstance.FindUser("invalidID")
				Expect(err).To(MatchError(errs.InvalidUserID))
			})
		})
	})

	Context("find many user", func() {
		userIDs := make([]string, 3)
		When("Success", func() {
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
				userIDs[2] = users[2].ID
			})
			It("should return user", func() {
				users, err := db.UserDBInstance.FindManyUser(userIDs)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				Expect(users[0]).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userIDs[0]),
					"FirstName":      Equal("testFirstName1"),
					"LastName":       Equal("testLastName1"),
					"Email":          Equal("testEmail1"),
					"HashedPassword": Equal("testHashPassword1"),
					"Role":           Equal(core.Student),
					"Courses":      Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
				Expect(users[1]).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userIDs[1]),
					"FirstName":      Equal("testFirstName2"),
					"LastName":       Equal("testLastName2"),
					"Email":          Equal("testEmail2"),
					"HashedPassword": Equal("testHashPassword2"),
					"Role":           Equal(core.Teacher),
					"Courses":      Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
				Expect(users[2]).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userIDs[2]),
					"FirstName":      Equal("testFirstName3"),
					"LastName":       Equal("testLastName3"),
					"Email":          Equal("testEmail3"),
					"HashedPassword": Equal("testHashPassword3"),
					"Role":           Equal(core.Admin),
					"Courses":      Equal([]string{}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
		When("Fail", func() {
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(2))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
			})
			It("should return user not found", func() {
				users, err := db.UserDBInstance.FindManyUser(userIDs)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(2))
			})
			It("should return invalid id", func() {
				_, err := db.UserDBInstance.FindManyUser([]string{"invalidID"})
				Expect(err).To(MatchError(errs.InvalidUserID))
			})
		})
	})

	Context("UpdateUser", func() {
		userIDs := make([]string, 3)
		When("Success", func() {
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
				userIDs[2] = users[2].ID
			})
			It("should update user", func() {
				user := core.User{
					ID:             userIDs[1],
					FirstName:      "testUpdatedFirstName2",
					LastName:       "testUpdatedLastName2",
					Email:          "testUpdatedEmail2",
					HashedPassword: "testUpdatedHashPassword2",
					Role:           core.Teacher,
					Courses:        []string{"testCurriculumID2"},
				}
				err := db.UserDBInstance.UpdateUser(&user)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":             Equal(userIDs[1]),
					"FirstName":      Equal("testUpdatedFirstName2"),
					"LastName":       Equal("testUpdatedLastName2"),
					"Email":          Equal("testUpdatedEmail2"),
					"HashedPassword": Equal("testUpdatedHashPassword2"),
					"Role":           Equal(core.Teacher),
					"Courses":      Equal([]string{"testCurriculumID2"}),
					"UpdatedAt":      Not(BeNil()),
				}))
			})
		})
	})

	Context("Delete user", func() {
		When("Success", func() {
			userIDs := make([]string, 3)
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHashPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
				userIDs[2] = users[2].ID
			})
			It("should delete user", func() {
				err := db.UserDBInstance.DeleteUser(userIDs[1])
				Expect(err).ShouldNot(HaveOccurred())
				user, err := db.UserDBInstance.FindUser(userIDs[1])
				Expect(err).To(MatchError(errs.UserNotFound))
				Expect(user).Should(BeZero())
			})
		})
	})

	Context("Delete many users", func() {
		When("Success", func() {
			userIDs := make([]string, 3)
			BeforeEach(func() {
				users := []core.User{
					{
						FirstName:      "testFirstName1",
						LastName:       "testLastName1",
						Email:          "testEmail1",
						HashedPassword: "testHashPassword1",
						Role:           core.Student,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName2",
						LastName:       "testLastName2",
						Email:          "testEmail2",
						HashedPassword: "testHashPassword2",
						Role:           core.Teacher,
						Courses:        []string{},
					},
					{
						FirstName:      "testFirstName3",
						LastName:       "testLastName3",
						Email:          "testEmail3",
						HashedPassword: "testHahPassword3",
						Role:           core.Admin,
						Courses:        []string{},
					},
				}
				err := db.UserDBInstance.InsertManyUser(users)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(3))
				userIDs[0] = users[0].ID
				userIDs[1] = users[1].ID
				userIDs[2] = users[2].ID
			})
			It("should delete many users", func() {
				err := db.UserDBInstance.DeleteManyUser(userIDs)
				Expect(err).ShouldNot(HaveOccurred())
				users, err := db.UserDBInstance.FindManyUser(userIDs)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users).Should(HaveLen(0))
			})
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
