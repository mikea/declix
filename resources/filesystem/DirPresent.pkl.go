// Code generated from Pkl module `mikea.declix.resources.FileSystem`. DO NOT EDIT.
package filesystem

type DirPresent interface {
	NodePresent
}

var _ DirPresent = (*DirPresentImpl)(nil)

type DirPresentImpl struct {
	Owner string `pkl:"owner"`

	Group string `pkl:"group"`

	Permissions string `pkl:"permissions"`
}

func (rcv *DirPresentImpl) GetOwner() string {
	return rcv.Owner
}

func (rcv *DirPresentImpl) GetGroup() string {
	return rcv.Group
}

func (rcv *DirPresentImpl) GetPermissions() string {
	return rcv.Permissions
}
