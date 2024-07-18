// Code generated from Pkl module `mikea.declix.resources.FileSystem`. DO NOT EDIT.
package filesystem

import "mikea/declix/resources"

type File interface {
	resources.Resource

	GetType() string

	GetPath() string

	GetState() any
}

var _ File = (*FileImpl)(nil)

type FileImpl struct {
	*resources.ResourceImpl

	Type string `pkl:"type"`

	Path string `pkl:"path"`

	State any `pkl:"state"`
}

func (rcv *FileImpl) GetType() string {
	return rcv.Type
}

func (rcv *FileImpl) GetPath() string {
	return rcv.Path
}

func (rcv *FileImpl) GetState() any {
	return rcv.State
}
