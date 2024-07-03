// Code generated from Pkl module `mikea.declix.System`. DO NOT EDIT.
package pkl

type Package interface {
	Resource

	GetType() string

	GetName() string

	GetStatus() string
}

var _ Package = (*PackageImpl)(nil)

type PackageImpl struct {
	*ResourceImpl

	Type string `pkl:"type"`

	Name string `pkl:"name"`

	Status string `pkl:"status"`
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
