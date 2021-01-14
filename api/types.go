package api

import (
	"encoding/json"

	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
)

type BasicRequest struct {
}

type BasicResponse struct {
	Err error `json:"error"`
}
type BasicError struct {
	Err *string `json:"error"`
}

type GroupRequest struct {
	Group *group.Group `json:"group"`
}
type GroupResponse struct {
	Group *group.Group `json:"group"`
	Err   *KmError     `json:"error"`
}

type GroupsResponse struct {
	Groups []group.Group `json:"groups"`
	Err    *KmError      `json:"error"`
}

type UserRequest struct {
	User       *passwd.User `json:"user"`
	CreateHome bool         `json:"create_home"`
}

type UserResponse struct {
	User *passwd.User `json:"user"`
	Err  *KmError     `json:"error"`
}

type UsersResponse struct {
	Users []passwd.User `json:"users"`
	Err   *KmError      `json:"error"`
}

// JSON marshaling of errors
// See: http://blog.magmalabs.io/2014/11/13/custom-error-marshaling-to-json-in-go.html
type KmError struct {
	error
}

func (me *KmError) MarshalJSON() ([]byte, error) {
	if me == nil {
		return []byte("null"), nil
	} else {
		return json.Marshal(me.Error())
	}
}

func (me *KmError) Err() error {
	if me == nil {
		return nil
	} else {
		return me
	}
}

func MakeKmError(e error) *KmError {
	return &KmError{e}
}
