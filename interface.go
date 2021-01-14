package sshkeymanager

import (
	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
)

type ClientInterface interface {
	GetGroups() ([]group.Group, error)
	ClearGroupCache() error
	FindGroup(group group.Group) (*group.Group, error)
	AddGroup(group group.Group) (*group.Group, error)
	DeleteGroup(group group.Group) (*group.Group, error)

	GetUsers() ([]passwd.User, error)
	ClearUserCache() error
	GetUserByUid(uid string) (*passwd.User, error)
	GetUserByName(name string) (*passwd.User, error)
	FindUser(user passwd.User) (*passwd.User, error)
	CreateHome(u passwd.User) (*passwd.User, error)
	AddUser(user passwd.User, createHome bool) (*passwd.User, error)
	DeleteUser(user passwd.User, removeHome bool) (*passwd.User, error)

	GetKeys(user passwd.User) ([]authorized_keys.SSHKey, error)
	FindKey(user passwd.User, key authorized_keys.SSHKey) (*authorized_keys.SSHKey, error)
	DeleteKey(user passwd.User, key authorized_keys.SSHKey) error
	AddKey(user passwd.User, key authorized_keys.SSHKey) error
	WriteKeys(user passwd.User, keys []authorized_keys.SSHKey) error
}
