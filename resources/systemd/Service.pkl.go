// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

type Service interface {
	Unit

	GetType() string

	GetState() *ServiceState

	GetId() string

	GetStateCmd() string

	GetCmds() *ServiceStateScripts
}

var _ Service = (*ServiceImpl)(nil)

type ServiceImpl struct {
	Type string `pkl:"type"`

	State *ServiceState `pkl:"state"`

	Id string `pkl:"id"`

	StateCmd string `pkl:"_stateCmd"`

	Cmds *ServiceStateScripts `pkl:"_cmds"`

	Name string `pkl:"name"`

	Systemctl string `pkl:"systemctl"`

	IsEnabled string `pkl:"_isEnabled"`

	IsActive string `pkl:"_isActive"`
}

func (rcv *ServiceImpl) GetType() string {
	return rcv.Type
}

func (rcv *ServiceImpl) GetState() *ServiceState {
	return rcv.State
}

func (rcv *ServiceImpl) GetId() string {
	return rcv.Id
}

func (rcv *ServiceImpl) GetStateCmd() string {
	return rcv.StateCmd
}

func (rcv *ServiceImpl) GetCmds() *ServiceStateScripts {
	return rcv.Cmds
}

func (rcv *ServiceImpl) GetName() string {
	return rcv.Name
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
