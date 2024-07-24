// Code generated from Pkl module `mikea.declix.resources.Users`. DO NOT EDIT.
package users

import "mikea/declix/resources"

type User interface {
	resources.Resource

	GetType() string

	GetLogin() string

	GetState() any

	GetId() string
}

var _ User = (*UserImpl)(nil)

type UserImpl struct {
	Type string `pkl:"type"`

	Login string `pkl:"login"`

	State any `pkl:"state"`

	Id string `pkl:"id"`
}

func (rcv *UserImpl) GetType() string {
	return rcv.Type
}

func (rcv *UserImpl) GetLogin() string {
	return rcv.Login
}

func (rcv *UserImpl) GetState() any {
	return rcv.State
}

func (rcv *UserImpl) GetId() string {
	return rcv.Id
}
