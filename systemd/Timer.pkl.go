// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type Timer interface {
	Unit

	GetType() string

	GetState() UnitState

	GetId() string
}

var _ Timer = (*TimerImpl)(nil)

type TimerImpl struct {
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

func (rcv *TimerImpl) GetType() string {
	return rcv.Type
}

func (rcv *TimerImpl) GetState() UnitState {
	return rcv.State
}

func (rcv *TimerImpl) GetId() string {
	return rcv.Id
}

func (rcv *TimerImpl) GetName() string {
	return rcv.Name
}

func (rcv *TimerImpl) GetUser() *string {
	return rcv.User
}

func (rcv *TimerImpl) GetFqName() string {
	return rcv.FqName
}

func (rcv *TimerImpl) GetSystemctl() string {
	return rcv.Systemctl
}

func (rcv *TimerImpl) GetIsEnabled() string {
	return rcv.IsEnabled
}

func (rcv *TimerImpl) GetIsActive() string {
	return rcv.IsActive
}

func (rcv *TimerImpl) GetStateCmd() string {
	return rcv.StateCmd
}

func (rcv *TimerImpl) GetCmds() *UnitStateScripts {
	return rcv.Cmds
}
