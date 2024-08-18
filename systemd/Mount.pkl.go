// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type Mount interface {
	Unit

	GetType() string

	GetState() UnitState

	GetId() string
}

var _ Mount = (*MountImpl)(nil)

type MountImpl struct {
	Type string `pkl:"type"`

	State UnitState `pkl:"state"`

	Id string `pkl:"id"`

	Name string `pkl:"name"`

	User *string `pkl:"user"`

	FqName string `pkl:"fqName"`

	Systemctl string `pkl:"systemctl"`

	IsEnabled string `pkl:"_isEnabled"`

	IsActive string `pkl:"_isActive"`

	StateCmd string `pkl:"_stateCmd"`

	Cmds *UnitStateScripts `pkl:"_cmds"`
}

func (rcv *MountImpl) GetType() string {
	return rcv.Type
}

func (rcv *MountImpl) GetState() UnitState {
	return rcv.State
}

func (rcv *MountImpl) GetId() string {
	return rcv.Id
}

func (rcv *MountImpl) GetName() string {
	return rcv.Name
}

func (rcv *MountImpl) GetUser() *string {
	return rcv.User
}

func (rcv *MountImpl) GetFqName() string {
	return rcv.FqName
}

func (rcv *MountImpl) GetSystemctl() string {
	return rcv.Systemctl
}

func (rcv *MountImpl) GetIsEnabled() string {
	return rcv.IsEnabled
}

func (rcv *MountImpl) GetIsActive() string {
	return rcv.IsActive
}

func (rcv *MountImpl) GetStateCmd() string {
	return rcv.StateCmd
}

func (rcv *MountImpl) GetCmds() *UnitStateScripts {
	return rcv.Cmds
}
