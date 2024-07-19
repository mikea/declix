// Code generated from Pkl module `mikea.declix.resources.apt`. DO NOT EDIT.
package apt

import (
	"mikea/declix/resources"
	"mikea/declix/resources/dpkg/state"
)

type Package interface {
	resources.Resource

	GetType() string

	GetName() string

	GetState() state.State

	GetUpdateBeforeInstall() bool
}

var _ Package = (*PackageImpl)(nil)

type PackageImpl struct {
	Type string `pkl:"type"`

	Name string `pkl:"name"`

	State state.State `pkl:"state"`

	UpdateBeforeInstall bool `pkl:"updateBeforeInstall"`
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

func (rcv *PackageImpl) GetUpdateBeforeInstall() bool {
	return rcv.UpdateBeforeInstall
}
