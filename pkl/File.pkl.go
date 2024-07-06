// Code generated from Pkl module `mikea.declix.System`. DO NOT EDIT.
package pkl

type File interface {
	Resource

	GetType() string

	GetPath() string

	GetContent() any

	GetOwner() string

	GetGroup() string

	GetPermissions() string
}

var _ File = (*FileImpl)(nil)

type FileImpl struct {
	*ResourceImpl

	Type string `pkl:"type"`

	Path string `pkl:"path"`

	Content any `pkl:"content"`

	Owner string `pkl:"owner"`

	Group string `pkl:"group"`

	Permissions string `pkl:"permissions"`
}

func (rcv *FileImpl) GetType() string {
	return rcv.Type
}

func (rcv *FileImpl) GetPath() string {
	return rcv.Path
}

func (rcv *FileImpl) GetContent() any {
	return rcv.Content
}

func (rcv *FileImpl) GetOwner() string {
	return rcv.Owner
}

func (rcv *FileImpl) GetGroup() string {
	return rcv.Group
}

func (rcv *FileImpl) GetPermissions() string {
	return rcv.Permissions
}
