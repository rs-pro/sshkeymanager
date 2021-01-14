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

	if req.User == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no user data provided")},
		})
		return
	}

	user, err := client.FindUser(*req.User)
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

func CreateHome(c *gin.Context) {
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

	if req.User == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no user data provided")},
		})
		return
	}

	user, err := client.CreateHome(*req.User)
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

	if req.User == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no user data provided")},
		})
		return
	}

	createHome := false
	if req.CreateHome != nil {
		createHome = *req.CreateHome
	}
	user, err := client.AddUser(*req.User, createHome)
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

func DeleteUser(c *gin.Context) {
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

	if req.User == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no user data provided")},
		})
		return
	}

	removeHome := false
	if req.RemoveHome != nil {
		removeHome = *req.RemoveHome
	}
	user, err := client.DeleteUser(*req.User, removeHome)
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
