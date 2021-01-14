package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetKeys(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := KeyRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, KeysResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	user, err := client.FindUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, KeysResponse{
			Err: &KmError{err},
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("user not found")},
		})
		return
	}

	keys, err := client.GetKeys(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, KeysResponse{
			Err: &KmError{err},
		})
		return
	}

	c.JSON(http.StatusOK, KeysResponse{
		User: user,
		Keys: keys,
	})
}

func FindKey(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := KeyRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, KeyResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	if req.Key == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no email or key data provided")},
		})
		return
	}

	user, err := client.FindUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{err},
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("user not found")},
		})
		return
	}

	key, err := client.FindKey(*user, *req.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{err},
		})
		return
	}

	c.JSON(http.StatusOK, KeyResponse{
		User: user,
		Key:  key,
	})
}

func AddKey(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := KeyRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, BasicResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	if req.Key == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no email or key data provided")},
		})
		return
	}

	user, err := client.FindUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BasicResponse{
			Err: &KmError{err},
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("user not found")},
		})
		return
	}

	err = client.AddKey(*user, *req.Key)

	if err != nil {
		c.JSON(http.StatusOK, BasicResponse{
			Err: &KmError{err},
		})
		return
	}
	c.JSON(http.StatusOK, BasicResponse{})
}

func DeleteKey(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := KeyRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, BasicResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	if req.Key == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("no email or key data provided")},
		})
		return
	}

	user, err := client.FindUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BasicResponse{
			Err: &KmError{err},
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("user not found")},
		})
		return
	}

	err = client.DeleteKey(*user, *req.Key)

	if err != nil {
		c.JSON(http.StatusOK, BasicResponse{
			Err: &KmError{err},
		})
		return
	}
	c.JSON(http.StatusOK, BasicResponse{})
}

func WriteKeys(c *gin.Context) {
	client := GetClient(c)
	if client == nil {
		return
	}

	req := KeysRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, BasicResponse{
			Err: &KmError{errors.Wrap(err, "bad json format")},
		})
		return
	}

	user, err := client.FindUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BasicResponse{
			Err: &KmError{err},
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, KeyResponse{
			Err: &KmError{errors.New("user not found")},
		})
		return
	}

	err = client.WriteKeys(*user, req.Keys)

	if err != nil {
		c.JSON(http.StatusOK, BasicResponse{
			Err: &KmError{err},
		})
		return
	}
	c.JSON(http.StatusOK, BasicResponse{})
}
