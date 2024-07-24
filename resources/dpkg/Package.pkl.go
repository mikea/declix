// Code generated from Pkl module `mikea.declix.resources.dpkg`. DO NOT EDIT.
package dpkg

import (
	"mikea/declix/resources"
	"mikea/declix/resources/dpkg/state"
)

type Package interface {
	resources.Resource

	GetType() string

	GetName() string

	GetState() state.State

	GetContent() any

	GetId() string
}

var _ Package = (*PackageImpl)(nil)

type PackageImpl struct {
	Type string `pkl:"type"`

	Name string `pkl:"name"`

	State state.State `pkl:"state"`

	Content any `pkl:"content"`

	Id string `pkl:"id"`
}

func (rcv *PackageImpl) GetType() string {
	return rcv.Type
}

func (rcv *PackageImpl) GetName() string {
	return rcv.Name
}

func (rcv *PackageImpl) GetState() state.State {
	return rcv.State
}

func (rcv *PackageImpl) GetContent() any {
	return rcv.Content
}

func (rcv *PackageImpl) GetId() string {
	return rcv.Id
}
