// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type TimerSection interface {
	Section

	GetName() any

	GetAccuracySec() *any

	GetFixedRandomDelay() *string

	GetOnActiveSec() *any

	GetOnBootSec() *any

	GetOnCalendar() *string

	GetOnClockChange() *string

	GetOnStartupSec() *any

	GetOnTimezoneChange() *string

	GetOnUnitActiveSec() *any

	GetOnUnitInactiveSec() *any

	GetPersistent() *string

	GetRandomizedDelaySec() *any

	GetRemainAfterElapse() *string

	GetUnit() *string

	GetWakeSystem() *string
}

var _ TimerSection = (*TimerSectionImpl)(nil)

type TimerSectionImpl struct {
	Name any `pkl:"name"`

	AccuracySec *any `pkl:"accuracySec"`

	FixedRandomDelay *string `pkl:"fixedRandomDelay"`

	OnActiveSec *any `pkl:"onActiveSec"`

	OnBootSec *any `pkl:"onBootSec"`

	OnCalendar *string `pkl:"onCalendar"`

	OnClockChange *string `pkl:"onClockChange"`

	OnStartupSec *any `pkl:"onStartupSec"`

	OnTimezoneChange *string `pkl:"onTimezoneChange"`

	OnUnitActiveSec *any `pkl:"onUnitActiveSec"`

	OnUnitInactiveSec *any `pkl:"onUnitInactiveSec"`

	Persistent *string `pkl:"persistent"`

	RandomizedDelaySec *any `pkl:"randomizedDelaySec"`

	RemainAfterElapse *string `pkl:"remainAfterElapse"`

	Unit *string `pkl:"unit"`

	WakeSystem *string `pkl:"wakeSystem"`
}

func (rcv *TimerSectionImpl) GetName() any {
	return rcv.Name
}

func (rcv *TimerSectionImpl) GetAccuracySec() *any {
	return rcv.AccuracySec
}

func (rcv *TimerSectionImpl) GetFixedRandomDelay() *string {
	return rcv.FixedRandomDelay
}

func (rcv *TimerSectionImpl) GetOnActiveSec() *any {
	return rcv.OnActiveSec
}

func (rcv *TimerSectionImpl) GetOnBootSec() *any {
	return rcv.OnBootSec
}

func (rcv *TimerSectionImpl) GetOnCalendar() *string {
	return rcv.OnCalendar
}

func (rcv *TimerSectionImpl) GetOnClockChange() *string {
	return rcv.OnClockChange
}

func (rcv *TimerSectionImpl) GetOnStartupSec() *any {
	return rcv.OnStartupSec
}

func (rcv *TimerSectionImpl) GetOnTimezoneChange() *string {
	return rcv.OnTimezoneChange
}

func (rcv *TimerSectionImpl) GetOnUnitActiveSec() *any {
	return rcv.OnUnitActiveSec
}

func (rcv *TimerSectionImpl) GetOnUnitInactiveSec() *any {
	return rcv.OnUnitInactiveSec
}

func (rcv *TimerSectionImpl) GetPersistent() *string {
	return rcv.Persistent
}

func (rcv *TimerSectionImpl) GetRandomizedDelaySec() *any {
	return rcv.RandomizedDelaySec
}

func (rcv *TimerSectionImpl) GetRemainAfterElapse() *string {
	return rcv.RemainAfterElapse
}

func (rcv *TimerSectionImpl) GetUnit() *string {
	return rcv.Unit
}

func (rcv *TimerSectionImpl) GetWakeSystem() *string {
	return rcv.WakeSystem
}
