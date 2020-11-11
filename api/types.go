package api

import (
	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
)

type BasicRequest struct {
}

type BasicResponse struct {
	Err error `json:"error"`
}

type GroupRequest struct {
	Group *group.Group `json:"group"`
}
type GroupResponse struct {
	Group *group.Group `json:"group"`
	Err   error        `json:"error"`
}

type GroupsResponse struct {
	Groups []group.Group `json:"groups"`
	Err    error         `json:"error"`
}

type UserRequest struct {
	User       *passwd.User `json:"user"`
	CreateHome bool         `json:"create_home"`
}

type UserResponse struct {
	User *passwd.User `json:"user"`
	Err  error        `json:"error"`
}

type UsersResponse struct {
	Users []passwd.User `json:"users"`
	Err   error         `json:"error"`
}
