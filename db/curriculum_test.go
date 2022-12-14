package db_test

// import (
// 	"time"

// 	"github.com/TendonT52/e-learning-tendon/db"
// 	"github.com/TendonT52/e-learning-tendon/internal/core"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("JwtToken", func() {
// 	BeforeEach(func() {
// 		db.NewClient("mongodb://admin:password@localhost:27017",
// 			db.MongoConfig{
// 				CreateTimeOut: time.Minute,
// 				FindTimeout:   time.Minute,
// 				UpdateTimeout: time.Minute,
// 				DeleteTimeout: time.Minute,
// 			})
// 		db.NewDB("tendon")
// 		db.NewUserDB("user_test")
// 		db.NewCurriculumDB("curriculum_test")
// 		db.CurriculumDBInstance.CleanUp()
// 	})

// 	Context("Insert curriculum to db", func() {
// 		var user core.User
// 		BeforeEach(func() {
// 			var err error
// 			user, err = db.UserDBInstance.InsertUser(
// 				"testFirstName",
// 				"testLastName",
// 				"testEmail",
// 				"testHashPassword",
// 				core.Student)
// 			Expect(err).ShouldNot(HaveOccurred())
// 			Expect(user).ShouldNot(BeZero())
// 		})

// 		When("Success", func() {
// 			It("should return curriculum", func() {
// 				curriculum, err := db.CurriculumDBInstance.InsertCurriculum(
// 					"testCurriculumName",
// 					"testCurriculumDescription",
// 					core.PrivateAccess,
// 					user.ID,
// 					[]string{"testLeasson1", "testLeasson2"},
// 				)
// 				Expect(err).ShouldNot(HaveOccurred())
// 				Expect(curriculum).ShouldNot(BeZero())
// 			})
// 		})
// 	})

// 	AfterEach(func() {
// 		db.CurriculumDBInstance.CleanUp()
// 		db.DisconnectMongo()
// 	})

// })
