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
		db.NewNodeDB("node_test")
		db.NodeDBInstance.Clear()
	})

	Context("test insert node", func() {
		It("should insert node", func() {
			node := core.Node{
				Type:     core.Text,
				Data:     "hello",
				CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
			}
			err := db.NodeDBInstance.InsertNode(&node)
			Expect(err).Should(BeNil())
			Expect(node).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Not(BeEmpty()),
				"Type":      Equal("text"),
				"Data":      Equal("hello"),
				"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
				"UpdatedAt": Not(BeNil()),
			}))
		})
	})

	Context("test insert many node", func() {
		It("should insert many node", func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(ConsistOf(
				MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeEmpty()),
					"Type":      Equal(core.Text),
					"Data":      Equal("hello"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeEmpty()),
					"Type":      Equal(core.Image),
					"Data":      Equal("path/to/image"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Not(BeEmpty()),
					"Type":      Equal(core.Pdf),
					"Data":      Equal("path/to/pdf"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
			))
		})
	})

	Context("test find node", func() {
		nodeIDs := make([]string, 3)
		BeforeEach(func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(HaveLen(3))
			nodeIDs[0] = nodes[0].ID
			nodeIDs[1] = nodes[1].ID
			nodeIDs[2] = nodes[2].ID
		})
		It("should find node", func() {
			foundNode, err := db.NodeDBInstance.FindNode(nodeIDs[1])
			Expect(err).Should(BeNil())
			Expect(foundNode).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(nodeIDs[1]),
				"Type":      Equal(core.Image),
				"Data":      Equal("path/to/image"),
				"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
				"UpdatedAt": Not(BeNil()),
			}))
		})
	})

	Context("test find many node", func() {
		nodeIDs := make([]string, 3)
		BeforeEach(func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(HaveLen(3))
			nodeIDs[0] = nodes[0].ID
			nodeIDs[1] = nodes[1].ID
			nodeIDs[2] = nodes[2].ID
		})
		It("should find many node", func() {
			foundNodes, err := db.NodeDBInstance.FindManyNode(nodeIDs)
			Expect(err).Should(BeNil())
			Expect(foundNodes).Should(ConsistOf(
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(nodeIDs[0]),
					"Type":      Equal(core.Text),
					"Data":      Equal("hello"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(nodeIDs[1]),
					"Type":      Equal(core.Image),
					"Data":      Equal("path/to/image"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
				MatchFields(IgnoreExtras, Fields{
					"ID":        Equal(nodeIDs[2]),
					"Type":      Equal(core.Pdf),
					"Data":      Equal("path/to/pdf"),
					"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
					"UpdatedAt": Not(BeNil()),
				}),
			))
		})
	})

	Context("test update node", func() {
		nodeIDs := make([]string, 3)
		BeforeEach(func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(HaveLen(3))
			nodeIDs[0] = nodes[0].ID
			nodeIDs[1] = nodes[1].ID
			nodeIDs[2] = nodes[2].ID
		})
		It("should update node", func() {
			node := core.Node{
				ID:       nodeIDs[1],
				Type:     core.Image,
				Data:     "path/to/image",
				CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
			}
			err := db.NodeDBInstance.UpdateNode(&node)
			Expect(err).Should(BeNil())
			Expect(node).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(nodeIDs[1]),
				"Type":      Equal(core.Image),
				"Data":      Equal("path/to/image"),
				"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
				"UpdatedAt": Not(BeNil()),
			}))
			foundNode, err := db.NodeDBInstance.FindNode(nodeIDs[1])
			Expect(err).Should(BeNil())
			Expect(foundNode).Should(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(nodeIDs[1]),
				"Type":      Equal(core.Image),
				"Data":      Equal("path/to/image"),
				"CreateBy":  Equal("5f9b2e7b2f2b0d8e5e6d7b5a"),
				"UpdatedAt": Not(BeNil()),
			}))
		})
	})

	Context("test delete node", func() {
		nodeIDs := make([]string, 3)
		BeforeEach(func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(HaveLen(3))
			nodeIDs[0] = nodes[0].ID
			nodeIDs[1] = nodes[1].ID
			nodeIDs[2] = nodes[2].ID
		})
		It("should delete node", func() {
			err := db.NodeDBInstance.DeleteNode(nodeIDs[0])
			Expect(err).Should(BeNil())
			node, err := db.NodeDBInstance.FindNode(nodeIDs[0])
			Expect(err).To(MatchError(errs.NodeNotFound))
			Expect(node).Should(BeZero())
		})
	})

	Context("test delete many node", func() {
		nodeIDs := make([]string, 3)
		BeforeEach(func() {
			nodes := []core.Node{
				{
					Type:     core.Text,
					Data:     "hello",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Image,
					Data:     "path/to/image",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
				{
					Type:     core.Pdf,
					Data:     "path/to/pdf",
					CreateBy: "5f9b2e7b2f2b0d8e5e6d7b5a",
				},
			}
			err := db.NodeDBInstance.InsertManyNode(nodes)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(HaveLen(3))
			nodeIDs[0] = nodes[0].ID
			nodeIDs[1] = nodes[1].ID
			nodeIDs[2] = nodes[2].ID
		})
		It("should delete many node", func() {
			err := db.NodeDBInstance.DeleteManyNode(nodeIDs)
			Expect(err).Should(BeNil())
			nodes, err := db.NodeDBInstance.FindManyNode(nodeIDs)
			Expect(err).Should(BeNil())
			Expect(nodes).Should(BeEmpty())
		})
	})

	AfterAll(func() {
		db.DisconnectMongo()
	})
})
