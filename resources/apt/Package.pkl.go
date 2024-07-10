// Code generated from Pkl module `mikea.declix.resources.apt`. DO NOT EDIT.
package apt

import "mikea/declix/resources"

type Package interface {
	resources.Resource

	GetType() string

	GetName() string

	GetStatus() string

	GetUpdateBeforeInstall() bool
}

var _ Package = (*PackageImpl)(nil)

type PackageImpl struct {
	*resources.ResourceImpl

	Type string `pkl:"type"`

	Name string `pkl:"name"`

	Status string `pkl:"status"`

	UpdateBeforeInstall bool `pkl:"updateBeforeInstall"`
}

func (rcv *PackageImpl) GetType() string {
	return rcv.Type
}

func (rcv *PackageImpl) GetName() string {
	return rcv.Name
}

func (rcv *PackageImpl) GetStatus() string {
	return rcv.Status
}

func (rcv *PackageImpl) GetUpdateBeforeInstall() bool {
	return rcv.UpdateBeforeInstall
}
