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

var _ = Describe("User", func() {
	BeforeEach(func() {
		db.NewClient("mongodb://admin:password@localhost:27017",
			db.MongoConfig{
				CreateTimeOut: time.Minute,
				FindTimeout:   time.Minute,
				UpdateTimeout: time.Minute,
				DeleteTimeout: time.Minute,
			})
		db.NewDB("tendon")
		db.NewUserDB("user_test")
		db.UserDBInstance.CleanUp()
	})

	Context("Insert user to db", func() {
		When("Success", func() {
			It("should return user", func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":               Not(BeNil()),
					"FirstName":        Equal("testFirstName"),
					"LastName":         Equal("testLastName"),
					"Email":            Equal("testEmail"),
					"HashPassword":     Equal("testHashPassword"),
					"Role":             Equal(core.Student),
					"UpdatedAt":        Not(BeNil()),
					"Curricula": BeEmpty(),
				}))
			})
		})
	})

	Context("Get user by email", func() {
		When("Success", func() {
			BeforeEach(func() {
				_, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return user", func() {
				user, err := db.UserDBInstance.GetUserByEmail(
					"testEmail",
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":               Not(BeNil()),
					"FirstName":        Equal("testFirstName"),
					"LastName":         Equal("testLastName"),
					"Email":            Equal("testEmail"),
					"HashPassword":     Equal("testHashPassword"),
					"Role":             Equal(core.Student),
					"UpdatedAt":        Not(BeNil()),
					"Curricula": BeEmpty(),
				}))
			})
		})
		When("Email not found", func() {
			It("should return error not found", func() {
				user, err := db.UserDBInstance.GetUserByEmail(
					"testEmail",
				)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(user).Should(BeZero())
			})
		})
	})

	Context("Get user by id", func() {
		var id string
		When("Success", func() {
			BeforeEach(func() {
				var err error
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				id = user.ID
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return user", func() {
				user, err := db.UserDBInstance.GetUserByID(
					id,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(MatchFields(IgnoreExtras, Fields{
					"ID":               Equal(id),
					"FirstName":        Equal("testFirstName"),
					"LastName":         Equal("testLastName"),
					"Email":            Equal("testEmail"),
					"HashPassword":     Equal("testHashPassword"),
					"Role":             Equal(core.Student),
					"UpdatedAt":        Not(BeNil()),
					"Curricula": BeEmpty(),
				}))
			})
		})
		When("ID not found", func() {
			It("should error not found", func() {
				user, err := db.UserDBInstance.GetUserByID(
					id,
				)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(user).Should(BeZero())
			})
		})
	})


	AfterEach(func() {
		db.UserDBInstance.CleanUp()
		db.DisconnectMongo()
	})
})
