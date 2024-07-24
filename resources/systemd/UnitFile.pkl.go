// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

import "mikea/declix/resources/filesystem"

type UnitFile interface {
	filesystem.File

	GetAfterAction() *AfterAction
}

var _ UnitFile = (*UnitFileImpl)(nil)

type UnitFileImpl struct {
	*filesystem.FileImpl

	AfterAction *AfterAction `pkl:"afterAction"`
}

func (rcv *UnitFileImpl) GetAfterAction() *AfterAction {
	return rcv.AfterAction
}
