// Code generated from Pkl module `mikea.declix.resources.FileSystem`. DO NOT EDIT.
package filesystem

type FilePresent interface {
	NodePresent

	GetContent() any
}

var _ FilePresent = (*FilePresentImpl)(nil)

type FilePresentImpl struct {
	Content any `pkl:"content"`

	Owner string `pkl:"owner"`

	Group string `pkl:"group"`

	Permissions string `pkl:"permissions"`
}

func (rcv *FilePresentImpl) GetContent() any {
	return rcv.Content
}

func (rcv *FilePresentImpl) GetOwner() string {
	return rcv.Owner
}

func (rcv *FilePresentImpl) GetGroup() string {
	return rcv.Group
}

func (rcv *FilePresentImpl) GetPermissions() string {
	return rcv.Permissions
}
