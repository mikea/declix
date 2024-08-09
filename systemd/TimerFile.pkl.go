// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

import "mikea/declix/content"

type TimerFile interface {
	content.Render

	GetUnit() UnitSection

	GetInstall() InstallSection

	GetTimer() TimerSection
}

var _ TimerFile = (*TimerFileImpl)(nil)

type TimerFileImpl struct {
	Unit UnitSection `pkl:"unit"`

	Install InstallSection `pkl:"install"`

	Timer TimerSection `pkl:"timer"`

	Result string `pkl:"_result"`

	Sha256 string `pkl:"_sha256"`
}

func (rcv *TimerFileImpl) GetUnit() UnitSection {
	return rcv.Unit
}

func (rcv *TimerFileImpl) GetInstall() InstallSection {
	return rcv.Install
}

func (rcv *TimerFileImpl) GetTimer() TimerSection {
	return rcv.Timer
}

func (rcv *TimerFileImpl) GetResult() string {
	return rcv.Result
}

func (rcv *TimerFileImpl) GetSha256() string {
	return rcv.Sha256
}
