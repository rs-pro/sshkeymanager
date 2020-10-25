package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs-pro/sshkeymanager/group"
)

type GetGroupsResponse struct {
	Groups []group.Group `json:"groups"`
	Err    error         `json:"error"`
}

func GetGroups(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}
	groups, err := client.GetGroups()
	c.JSON(http.StatusOK, GetGroupsResponse{
		Groups: groups,
		Err:    err,
	})
}

type FindGroupRequest struct {
	Group *group.Group `json:"group"`
}
type FindGroupResponse struct {
	Group *group.Group `json:"group"`
	Err   error        `json:"error"`
}

func FindGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := FindGroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, FindGroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	g, err := client.AddGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	log.Println("added group:", g.GID, g.Name, g.Members)

	c.JSON(http.StatusOK, AddGroupResponse{Group: g})
}

type AddGroupRequest struct {
	Group *group.Group `json:"group"`
}
type AddGroupResponse struct {
	Group *group.Group `json:"group"`
	Err   error        `json:"error"`
}

func AddGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := AddGroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, AddGroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	g, err := client.AddGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AddGroupResponse{
			Err: err,
		})
		return
	}
	log.Println("added group:", g.GID, g.Name, g.Members)

	c.JSON(http.StatusOK, AddGroupResponse{Group: g})
}

type DeleteGroupRequest struct {
	Group *group.Group `json:"group"`
}
type DeleteGroupResponse struct {
	Group *group.Group `json:"group"`
	Err   error        `json:"error"`
}

func DeleteGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := DeleteGroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, AddGroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	if req.Group.Name == "" && req.Group.GID != "" {
		req.Group, _ = client.FindGroup(req.Group)
	}

	_, err = client.DeleteGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AddGroupResponse{
			Err: err,
		})
		return
	}

	c.JSON(http.StatusOK, DeleteGroupResponse{Group: req.Group})
}
