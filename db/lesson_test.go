package db_test

import "github.com/TendonT52/e-learning-tendon/db"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("node", func() {
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
		db.NewLessonDB("lesson_test")
		db.NewNodeDB("node_test")
		db.NodeDBInstance.CleanUp()
		db.UserDBInstance.CleanUp()
	})

	context("Insert lesson to db", func() {
	})

	afterEach(func() {
		db.NodeDBInstance.CleanUp()
		db.UserDBInstance.CleanUp()
		db.LessonDBInstance.CleanUp()
	})
})