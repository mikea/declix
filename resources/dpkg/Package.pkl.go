// Code generated from Pkl module `mikea.declix.resources.dpkg`. DO NOT EDIT.
package dpkg

import "mikea/declix/resources"

type Package interface {
	resources.Resource

	GetType() string

	GetName() string

	GetStatus() string

	GetUrl() string
}

var _ Package = (*PackageImpl)(nil)

type PackageImpl struct {
	*resources.ResourceImpl

	Type string `pkl:"type"`

	Name string `pkl:"name"`

	Status string `pkl:"status"`

	Url string `pkl:"url"`
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

func (rcv *PackageImpl) GetUrl() string {
	return rcv.Url
}
