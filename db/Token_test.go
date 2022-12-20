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
				InsertTimeOut: time.Minute,
				FindTimeOut:   time.Minute,
				UpdateTimeOut: time.Minute,
				DeleteTimeOut: time.Minute,
			})
		db.NewDB("tendon")
		db.NewJwtDB("jwtToken_test")
		db.JwtDBInstance.Clear()
	})

	Context("Insert jwt token to db", func() {
		When("Success", func() {
			It("should return jwt token", func() {
				tokenString, err := db.JwtDBInstance.InsertJwt(time.Now().Add(time.Minute))
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
				tokenString, err = db.JwtDBInstance.InsertJwt(time.Now().Add(time.Minute))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tokenString).ShouldNot(BeZero())
			})
			It("should return nil", func() {
				err := db.JwtDBInstance.CheckJwt(tokenString)
				Expect(err).Should(BeNil())
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				err := db.JwtDBInstance.CheckJwt(tokenString)
				Expect(err).To(MatchError(errs.TokenNotfound))
			})
		})
	})

	Context("Delete jwt token", func() {
		var tokenString string
		When("Success", func() {
			BeforeEach(func() {
				var err error
				tokenString, err = db.JwtDBInstance.InsertJwt(time.Now().Add(time.Minute))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tokenString).ShouldNot(BeZero())
			})
			It("should return nil", func() {
				err := db.JwtDBInstance.DeleteJwt(tokenString)
				Expect(err).Should(BeNil())
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				err := db.JwtDBInstance.DeleteJwt(tokenString)
				Expect(err).To(MatchError(errs.TokenNotfound))
			})
		})
	})

	AfterEach(func() {
		db.JwtDBInstance.Clear()
		db.DisconnectMongo()
	})
})
