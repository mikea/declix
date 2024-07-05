// Code generated from Pkl module `mikea.declix.System`. DO NOT EDIT.
package pkl

type File interface {
	Resource

	GetType() string

	GetPath() string

	GetContentFile() string

	GetOwner() string

	GetGroup() string

	GetPermissions() string
}

var _ File = (*FileImpl)(nil)

type FileImpl struct {
	*ResourceImpl

	Type string `pkl:"type"`

	Path string `pkl:"path"`

	ContentFile string `pkl:"contentFile"`

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

func (rcv *FileImpl) GetContentFile() string {
	return rcv.ContentFile
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
