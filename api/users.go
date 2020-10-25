package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs-pro/sshkeymanager/passwd"
)

type GetUsersResponse struct {
	Users []passwd.User `json:"users"`
	Err   error         `json:"error"`
}

func GetUsers(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}
	users, err := client.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetUsersResponse{
			Err: err,
		})
		return
	}

	c.JSON(http.StatusOK, GetUsersResponse{
		Users: users,
	})
}

type AddUserRequest struct {
	User       *passwd.User `json:"user"`
	CreateHome bool         `json:"create_home"`
}

type AddUserResponse struct {
	User *passwd.User `json:"user"`
	Err  error        `json:"error"`
}

func AddUser(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := AddUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, AddUserResponse{
			Err: errors.Wrap(err, "bad json format"),
		})
		return
	}

	user, err := client.AddUser(req.User, req.CreateHome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AddUserResponse{
			Err: err,
		})
		return
	}

	c.JSON(http.StatusOK, AddUserResponse{
		User: user,
	})
}
