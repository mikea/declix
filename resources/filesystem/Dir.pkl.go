// Code generated from Pkl module `mikea.declix.resources.FileSystem`. DO NOT EDIT.
package filesystem

type Dir interface {
	Node

	GetType() string

	GetState() any

	GetId() string

	GetDetermineStateCmd() string
}

var _ Dir = (*DirImpl)(nil)

type DirImpl struct {
	Type string `pkl:"type"`

	State any `pkl:"state"`

	Id string `pkl:"id"`

	DetermineStateCmd string `pkl:"_determineStateCmd"`

	Path string `pkl:"path"`
}

func (rcv *DirImpl) GetType() string {
	return rcv.Type
}

func (rcv *DirImpl) GetState() any {
	return rcv.State
}

func (rcv *DirImpl) GetId() string {
	return rcv.Id
}

func (rcv *DirImpl) GetDetermineStateCmd() string {
	return rcv.DetermineStateCmd
}

func (rcv *DirImpl) GetPath() string {
	return rcv.Path
}
