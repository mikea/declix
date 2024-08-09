// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

import "github.com/apple/pkl-go/pkl"

func init() {
	pkl.RegisterMapping("mikea.declix.systemd#UnitFile", UnitFileImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#AfterAction", AfterAction{})
	pkl.RegisterMapping("mikea.declix.systemd#_UnitStateScripts", UnitStateScripts{})
	pkl.RegisterMapping("mikea.declix.systemd#UnitState", UnitStateImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#Service", ServiceImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#Socket", SocketImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#UnitSection", UnitSectionImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#InstallSection", InstallSectionImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#ServiceSection", ServiceSectionImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#ServiceFile", ServiceFileImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#TimerSection", TimerSectionImpl{})
	pkl.RegisterMapping("mikea.declix.systemd#TimerFile", TimerFileImpl{})
}
