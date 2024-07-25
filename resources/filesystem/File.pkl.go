// Code generated from Pkl module `mikea.declix.resources.FileSystem`. DO NOT EDIT.
package filesystem

type File interface {
	Node

	GetType() string

	GetState() any

	GetId() string

	GetDetermineStateCmd() string
}

var _ File = (*FileImpl)(nil)

type FileImpl struct {
	Type string `pkl:"type"`

	State any `pkl:"state"`

	Id string `pkl:"id"`

	DetermineStateCmd string `pkl:"_determineStateCmd"`

	Path string `pkl:"path"`
}

func (rcv *FileImpl) GetType() string {
	return rcv.Type
}

func (rcv *FileImpl) GetState() any {
	return rcv.State
}

func (rcv *FileImpl) GetId() string {
	return rcv.Id
}

func (rcv *FileImpl) GetDetermineStateCmd() string {
	return rcv.DetermineStateCmd
}

func (rcv *FileImpl) GetPath() string {
	return rcv.Path
}
