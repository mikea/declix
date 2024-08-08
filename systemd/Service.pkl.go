// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type Service interface {
	Unit

	GetType() string

	GetState() UnitState

	GetId() string
}

var _ Service = (*ServiceImpl)(nil)

type ServiceImpl struct {
	Type string `pkl:"type"`

	State UnitState `pkl:"state"`

	Id string `pkl:"id"`

	Name string `pkl:"name"`

	User *string `pkl:"user"`

	Systemctl string `pkl:"systemctl"`

	IsEnabled string `pkl:"_isEnabled"`

	IsActive string `pkl:"_isActive"`

	StateCmd string `pkl:"_stateCmd"`

	Cmds *UnitStateScripts `pkl:"_cmds"`
}

func (rcv *ServiceImpl) GetType() string {
	return rcv.Type
}

func (rcv *ServiceImpl) GetState() UnitState {
	return rcv.State
}

func (rcv *ServiceImpl) GetId() string {
	return rcv.Id
}

func (rcv *ServiceImpl) GetName() string {
	return rcv.Name
}

func (rcv *ServiceImpl) GetUser() *string {
	return rcv.User
}

func (rcv *ServiceImpl) GetSystemctl() string {
	return rcv.Systemctl
}

func (rcv *ServiceImpl) GetIsEnabled() string {
	return rcv.IsEnabled
}

func (rcv *ServiceImpl) GetIsActive() string {
	return rcv.IsActive
}

func (rcv *ServiceImpl) GetStateCmd() string {
	return rcv.StateCmd
}

func (rcv *ServiceImpl) GetCmds() *UnitStateScripts {
	return rcv.Cmds
}
