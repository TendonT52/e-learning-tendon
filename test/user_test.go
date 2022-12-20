package test

import (
	"net/http"

	"github.com/TendonT52/e-learning-tendon/config"
	"github.com/TendonT52/e-learning-tendon/db"
	"github.com/TendonT52/e-learning-tendon/handlers"
	. "github.com/onsi/ginkgo/v2"

	// . "github.com/onsi/gomega"
	// . "github.com/onsi/gomega/gstruct"
	"github.com/steinfletcher/apitest"
)

var _ = Describe("test course db", Ordered, func() {
	var t GinkgoTInterface

	BeforeAll(func() {
		config.LoadConfigTest()
		config.SetupInstance()
		handlers.SetupRouter()
		db.UserDBInstance.Clear()
		db.JwtDBInstance.Clear()
		db.CourseDBInstance.Clear()
		db.LessonDBInstance.Clear()
		db.NodeDBInstance.Clear()
		t = GinkgoT()
	})

	Context("Successful CookieMatching", func() {
		It("cookies should be set correctly", func() {
			apitest.New().
				Handler(handlers.Router).
				Post("/api/v1/user/sign-up").
				JSON(`
				{
					"firstName": "firstnametest1",	
					"lastName": "lastnametest1",
					"email": "email@test.com",
					"password": "password1"
				}
				`).
				Expect(t).
				CookiePresent("access_token").
				CookiePresent("refresh_token").
				Status(http.StatusCreated).
				End()
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
