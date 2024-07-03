// Code generated from Pkl module `mikea.declix.System`. DO NOT EDIT.
package pkl

type File interface {
	Resource

	GetType() string

	GetPath() string
}

var _ File = (*FileImpl)(nil)

type FileImpl struct {
	*ResourceImpl

	Type string `pkl:"type"`

	Path string `pkl:"path"`
}

func (rcv *FileImpl) GetType() string {
	return rcv.Type
}

func (rcv *FileImpl) GetPath() string {
	return rcv.Path
}
