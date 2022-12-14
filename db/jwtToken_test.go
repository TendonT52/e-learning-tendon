package db_test

import (
	"time"

	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jwt", func() {
	BeforeEach(func() {
		db.NewClient("mongodb://admin:password@localhost:27017",
			db.MongoConfig{
				CreateTimeOut: time.Minute,
				FindTimeout:   time.Minute,
				UpdateTimeout: time.Minute,
				DeleteTimeout: time.Minute,
			})
		db.NewDB("tendon")
		db.NewJwtDB("jwtToken_test")
		db.JwtDBInstance.CleanUp()
	})

	Context("Insert jwt token to db", func() {
		When("Success", func() {
			It("should return jwt token", func() {
				tokenString, err := db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tokenString).ShouldNot(BeZero())
			})
		})
	})

	Context("Check jwt token", func() {
		var tokenString string
		When("Success", func() {
			BeforeEach(func() {
				var err error
				tokenString, err = db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tokenString).ShouldNot(BeZero())
			})
			It("should return nil", func() {
				err := db.JwtDBInstance.CheckJwtToken(tokenString)
				Expect(err).Should(BeNil())
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				err := db.JwtDBInstance.CheckJwtToken(tokenString)
				Expect(err).To(MatchError(errs.ErrNotFound))
			})
		})
	})

	Context("Delete jwt token", func() {
		var tokenString string
		When("Success", func() {
			BeforeEach(func() {
				var err error
				tokenString, err = db.JwtDBInstance.InsertJwtToken(time.Now().Add(time.Minute))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tokenString).ShouldNot(BeZero())
			})
			It("should return nil", func() {
				err := db.JwtDBInstance.DeleteJwtToken(tokenString)
				Expect(err).Should(BeNil())
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				err := db.JwtDBInstance.DeleteJwtToken(tokenString)
				Expect(err).To(MatchError(errs.ErrNotFound))
			})
		})
	})

	AfterEach(func() {
		db.JwtDBInstance.CleanUp()
		db.DisconnectMongo()
	})
})
