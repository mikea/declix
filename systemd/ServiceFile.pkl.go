// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

import "mikea/declix/content"

type ServiceFile interface {
	content.Render

	GetUnit() UnitSection

	GetInstall() InstallSection

	GetService() ServiceSection
}

var _ ServiceFile = (*ServiceFileImpl)(nil)

type ServiceFileImpl struct {
	Unit UnitSection `pkl:"unit"`

	Install InstallSection `pkl:"install"`

	Service ServiceSection `pkl:"service"`

	Result string `pkl:"_result"`

	Sha256 string `pkl:"_sha256"`
}

func (rcv *ServiceFileImpl) GetUnit() UnitSection {
	return rcv.Unit
}

func (rcv *ServiceFileImpl) GetInstall() InstallSection {
	return rcv.Install
}

func (rcv *ServiceFileImpl) GetService() ServiceSection {
	return rcv.Service
}

func (rcv *ServiceFileImpl) GetResult() string {
	return rcv.Result
}

func (rcv *ServiceFileImpl) GetSha256() string {
	return rcv.Sha256
}
