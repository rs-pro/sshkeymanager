package api

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetUsers(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}
	users, err := client.GetUsers()
	if err != nil {
		spew.Dump(err)
		c.JSON(http.StatusInternalServerError, UsersResponse{
			Err: &KmError{err},
		})
		return
	}

	c.JSON(http.StatusOK, UsersResponse{
		Users: users,
	})
}

func FindUser(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, UserResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	user, err := client.AddUser(req.User, req.CreateHome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Err: &KmError{err},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		User: user,
	})
}

func AddUser(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, UserResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	user, err := client.AddUser(req.User, req.CreateHome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Err: &KmError{err},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		User: user,
	})
}
