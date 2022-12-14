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
		db.NewNodeDB("node_test")
		db.NewUserDB("user_test")
		db.NodeDBInstance.CleanUp()
		db.UserDBInstance.CleanUp()
	})

	Context("Insert node to db", func() {
		When("Success", func() {
			var userID string
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID
			})
			It("should return node", func() {
				node, err := db.NodeDBInstance.InsertNode(
					core.Image,
					"path/to/image",
					userID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Image),
					"Data":      Equal("path/to/image"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				node, err := db.NodeDBInstance.InsertNode(
					core.Image,
					"path/to/image",
					"invalidUserID",
				)
				Expect(err).To(MatchError(errs.ErrWrongFormat))
				Expect(node).To(BeZero())
			})
		})
	})

	Context("Insert node many to db", func() {
		var userID string
		When("Success", func() {
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID
			})
			It("should return node", func() {
				nodes, err := db.NodeDBInstance.InsertNodeMany(
					[]string{core.Image, core.Video, core.Text, core.Pdf, core.Sound},
					[]string{"path/to/image", "path/to/video", "path/to/text", "path/to/pdf", "path/to/sound"},
					userID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(nodes).To(HaveLen(5))
				Expect(nodes[0]).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Image),
					"Data":      Equal("path/to/image"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
				Expect(nodes[1]).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Video),
					"Data":      Equal("path/to/video"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
				Expect(nodes[2]).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Text),
					"Data":      Equal("path/to/text"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
				Expect(nodes[3]).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Pdf),
					"Data":      Equal("path/to/pdf"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
				Expect(nodes[4]).To(MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeNil()),
					"Type":      Equal(core.Sound),
					"Data":      Equal("path/to/sound"),
					"CreateBy":  Equal(userID),
					"UpdatedAt": Not(BeNil()),
				}))
			})
		})
		When("Fail", func() {
			It("should return error", func() {
				nodes, err := db.NodeDBInstance.InsertNodeMany(
					[]string{core.Video, core.Text, core.Pdf, core.Sound},
					[]string{"path/to/image", "path/to/video", "path/to/text", "path/to/pdf", "path/to/sound"},
					userID,
				)
				Expect(err).To(MatchError(errs.ErrWrongFormat))
				Expect(nodes).To(BeZero())
			})
		})
	})

	Context("Get node from db by id", func() {
		var nodeID string
		var userID string
		When("Success", func() {
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID

				node, err := db.NodeDBInstance.InsertNode(
					core.Image,
					"path/to/image",
					user.ID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node).ShouldNot(BeZero())
				nodeID = node.ID
			})
			It("should return node", func() {
				node, err := db.NodeDBInstance.GetNodeByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node).To(MatchFields(IgnoreExtras, Fields{
					"ID":       Equal(nodeID),
					"Type":     Equal(core.Image),
					"Data":     Equal("path/to/image"),
					"CreateBy": Equal(userID),
				}))
			})
		})
		When("Fail", func() {
			It("should return wrong format error", func() {
				node, err := db.NodeDBInstance.GetNodeByID("invalidNodeID")
				Expect(err).To(MatchError(errs.ErrWrongFormat))
				Expect(node).To(BeZero())
			})
			It("should return not found error", func() {
				node, err := db.NodeDBInstance.GetNodeByID(nodeID)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(node).To(BeZero())
			})
		})
	})

	Context("Get node many from db by id", func() {
		var nodeID []string
		var userID string
		When("Success", func() {
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID

				nodes, err := db.NodeDBInstance.InsertNodeMany(
					[]string{core.Image, core.Video, core.Text, core.Pdf, core.Sound},
					[]string{"path/to/image", "path/to/video", "path/to/text", "path/to/pdf", "path/to/sound"},
					userID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(nodes).To(HaveLen(5))
				nodeID = []string{nodes[0].ID, nodes[1].ID, nodes[2].ID, nodes[3].ID, nodes[4].ID}
			})
			It("should return node", func() {
				nodes, err := db.NodeDBInstance.GetNodeManyByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(nodes).To(HaveLen(5))
			})
		})
		When("Fail", func() {
			It("should return wrong format error", func() {
				nodes, err := db.NodeDBInstance.GetNodeManyByID([]string{"invalidNodeID"})
				Expect(err).To(MatchError(errs.ErrWrongFormat))
				Expect(nodes).To(BeZero())
			})
			It("should return not found error", func() {
				nodes, err := db.NodeDBInstance.GetNodeManyByID(nodeID)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(nodes).To(BeZero())
			})
		})
	})

	Context("Delete node from db by id", func() {
		var nodeID string
		var userID string
		When("Success", func() {
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID

				node, err := db.NodeDBInstance.InsertNode(
					core.Image,
					"path/to/image",
					userID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node).ShouldNot(BeZero())
				nodeID = node.ID
			})
			It("should return node", func() {
				err := db.NodeDBInstance.DeleteNodeByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				node, err := db.NodeDBInstance.GetNodeByID(nodeID)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(node).To(BeZero())
			})
		})
	})

	Context("Delete node many from db by id", func() {
		var nodeID []string
		var userID string
		When("Success", func() {
			BeforeEach(func() {
				user, err := db.UserDBInstance.InsertUser(
					"testFirstName",
					"testLastName",
					"testEmail",
					"testHashPassword",
					core.Student)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).ShouldNot(BeZero())
				userID = user.ID

				nodes, err := db.NodeDBInstance.InsertNodeMany(
					[]string{core.Image, core.Video, core.Text, core.Pdf, core.Sound},
					[]string{"path/to/image", "path/to/video", "path/to/text", "path/to/pdf", "path/to/sound"},
					userID,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(nodes).To(HaveLen(5))
				nodeID = []string{nodes[0].ID, nodes[1].ID, nodes[2].ID, nodes[3].ID, nodes[4].ID}
			})
			It("should return node", func() {
				err := db.NodeDBInstance.DeleteNodeManyByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				nodes, err := db.NodeDBInstance.GetNodeManyByID(nodeID)
				Expect(err).To(MatchError(errs.ErrNotFound))
				Expect(nodes).To(BeZero())
			})

		})
		When("Fail", func() {
			It("should return wrong format error", func() {
				err := db.NodeDBInstance.DeleteNodeManyByID([]string{"invalidNodeID"})
				Expect(err).To(MatchError(errs.ErrWrongFormat))
			})
			It("should return not found error", func() {
				err := db.NodeDBInstance.DeleteNodeManyByID(nodeID)
				Expect(err).To(MatchError(errs.ErrNotFound))
			})
		})
	})
	AfterEach(func() {
		db.UserDBInstance.CleanUp()
		db.NodeDBInstance.CleanUp()
		db.DisconnectMongo()
	})
})
