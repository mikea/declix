// Code generated from Pkl module `mikea.declix.resources.Users`. DO NOT EDIT.
package users

import "mikea/declix/resources"

type Group interface {
	resources.Resource

	GetType() string

	GetName() string

	GetState() any

	GetId() string
}

var _ Group = (*GroupImpl)(nil)

type GroupImpl struct {
	Type string `pkl:"type"`

	Name string `pkl:"name"`

	State any `pkl:"state"`

	Id string `pkl:"id"`
}

func (rcv *GroupImpl) GetType() string {
	return rcv.Type
}

func (rcv *GroupImpl) GetName() string {
	return rcv.Name
}

func (rcv *GroupImpl) GetState() any {
	return rcv.State
}

func (rcv *GroupImpl) GetId() string {
	return rcv.Id
}
