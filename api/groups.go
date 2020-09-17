package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager/group"
)

func ListGroups(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}
	users, err := client.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func AddGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	g := &group.Group{}
	err := c.BindJSON(g)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "bad json format: " + err.Error(),
		})
		return
	}

	g, err = client.AddGroup(g)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	log.Println("added group:", g.GID, g.Name, g.Members)

	c.JSON(http.StatusOK, g)
}

func DeleteGroup(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	g := &group.Group{}
	err := c.BindJSON(g)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "bad json format: " + err.Error(),
		})
		return
	}

	if g.Name == "" && g.GID != "" {
		g = client.FindGroup(g)
	}

	_, err = client.DeleteGroup(g)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	log.Println("deleted group:", g.GID, g.Name)

	c.JSON(http.StatusOK, g)
}
