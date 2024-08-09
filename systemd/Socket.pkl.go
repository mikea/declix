// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type Socket interface {
	Unit

	GetType() string

	GetState() UnitState

	GetId() string
}

var _ Socket = (*SocketImpl)(nil)

type SocketImpl struct {
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

func (rcv *SocketImpl) GetType() string {
	return rcv.Type
}

func (rcv *SocketImpl) GetState() UnitState {
	return rcv.State
}

func (rcv *SocketImpl) GetId() string {
	return rcv.Id
}

func (rcv *SocketImpl) GetName() string {
	return rcv.Name
}

func (rcv *SocketImpl) GetUser() *string {
	return rcv.User
}

func (rcv *SocketImpl) GetFqName() string {
	return rcv.FqName
}

func (rcv *SocketImpl) GetSystemctl() string {
	return rcv.Systemctl
}

func (rcv *SocketImpl) GetIsEnabled() string {
	return rcv.IsEnabled
}

func (rcv *SocketImpl) GetIsActive() string {
	return rcv.IsActive
}

func (rcv *SocketImpl) GetStateCmd() string {
	return rcv.StateCmd
}

func (rcv *SocketImpl) GetCmds() *UnitStateScripts {
	return rcv.Cmds
}
