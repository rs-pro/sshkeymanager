package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetGroups(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}
	groups, err := client.GetGroups()

	if err != nil {
		c.JSON(http.StatusInternalServerError, GroupsResponse{
			Err: err,
		})
		return
	}

	c.JSON(http.StatusOK, GroupsResponse{
		Groups: groups,
		Err:    err,
	})
}

func FindGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := GroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, GroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	g, err := client.AddGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GroupResponse{
			Err: err,
		})
		return
	}
	log.Println("added group:", g.GID, g.Name, g.Members)

	c.JSON(http.StatusOK, GroupResponse{Group: g})
}

func AddGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := GroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, GroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	g, err := client.AddGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GroupResponse{
			Err: err,
		})
		return
	}
	log.Println("added group:", g.GID, g.Name, g.Members)

	c.JSON(http.StatusOK, GroupResponse{Group: g})
}

func DeleteGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := GroupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, GroupResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	if req.Group.Name == "" && req.Group.GID != "" {
		req.Group, _ = client.FindGroup(req.Group)
	}

	_, err = client.DeleteGroup(req.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GroupResponse{
			Err: err,
		})
		return
	}

	c.JSON(http.StatusOK, GroupResponse{Group: req.Group})
}
