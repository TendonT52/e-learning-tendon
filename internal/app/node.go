package app

import (
	"errors"
	"net/http"

	"github.com/TendonT52/e-learning-tendon/internal/core"
	"github.com/TendonT52/e-learning-tendon/internal/pkg/errs"
)

func CreateNode(node *core.Node) error {
	_, err := reposInstance.UserDB.FindUser(node.CreateBy)
	if err != nil {
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"create by id is not invalid",
			)
		}
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"create by id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	err = reposInstance.NodeDB.InsertNode(node)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't insert node to database",
		)
	}
	return nil
}

func GetNode(nodeId string) (core.Node,	error){
	node, err := reposInstance.NodeDB.FindNode(nodeId)
	if err != nil {
		if errors.Is(err, errs.InvalidNodeID) {
			return core.Node{}, errs.NewHttpError(
				http.StatusBadRequest,
				"node id is not invalid",
			)
		}
		if errors.Is(err, errs.NodeNotFound) {
			return core.Node{}, errs.NewHttpError(
				http.StatusNotFound,
				"node id is not found",
			)
		}
		return core.Node{}, errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find node in database",
		)
	}
	return node, nil
}

func UpdateNode(node *core.Node) error {
	_, err := reposInstance.UserDB.FindUser(node.CreateBy)
	if err != nil {
		if errors.Is(err, errs.InvalidUserID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"update by id is not invalid",
			)
		}
		if errors.Is(err, errs.UserNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"update by id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't find user in database",
		)
	}
	err = reposInstance.NodeDB.UpdateNode(node)
	if err != nil {
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't update node in database",
		)
	}
	return nil
}

func DeleteNode(nodeId string) error {
	err := reposInstance.NodeDB.DeleteNode(nodeId)
	if err != nil {
		if errors.Is(err, errs.InvalidNodeID) {
			return errs.NewHttpError(
				http.StatusBadRequest,
				"node id is not invalid",
			)
		}
		if errors.Is(err, errs.NodeNotFound) {
			return errs.NewHttpError(
				http.StatusNotFound,
				"node id is not found",
			)
		}
		return errs.NewHttpError(
			http.StatusInternalServerError,
			"can't delete node in database",
		)
	}
	return nil
}