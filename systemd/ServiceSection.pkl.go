// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

import (
	"mikea/declix/systemd/serviceexittype"
	"mikea/declix/systemd/servicetype"
)

type ServiceSection interface {
	Section

	GetName() any

	GetType() *servicetype.ServiceType

	GetBusName() *string

	GetEnvironment() map[string]string

	GetEnvironmentFile() *string

	GetExecCondition() *string

	GetExecReload() *string

	GetExecStart() *string

	GetExecStartPost() *string

	GetExecStartPre() *string

	GetExecStop() *string

	GetExecStopPost() *string

	GetExitType() *serviceexittype.ServiceExitType

	GetFileDescriptorStoreMax() *string

	GetFileDescriptorStorePreserve() *string

	GetGuessMainPID() *string

	GetNonBlocking() *string

	GetNotifyAccess() *string

	GetOOMPolicy() *string

	GetOpenFile() *string

	GetPIDFile() *string

	GetReloadSignal() *string

	GetRemainAfterExit() *string

	GetRestart() *string

	GetRestartForceExitStatus() *string

	GetRestartMaxDelaySec() *string

	GetRestartMode() *string

	GetRestartPreventExitStatus() *string

	GetRestartSec() *string

	GetRestartSteps() *string

	GetRootDirectoryStartOnly() *string

	GetRuntimeMaxSec() *string

	GetRuntimeRandomizedExtraSec() *string

	GetSockets() *string

	GetSuccessExitStatus() *string

	GetTimeoutAbortSec() *uint

	GetTimeoutSec() *uint

	GetTimeoutStartFailureMode() *uint

	GetTimeoutStartSec() *uint

	GetTimeoutStopFailureMode() *uint

	GetTimeoutStopSec() *uint

	GetUSBFunctionDescriptors() *string

	GetUSBFunctionStrings() *string

	GetWatchdogSec() *string

	GetFinalKillSignal() *string

	GetKillMode() *string

	GetKillSignal() *string

	GetRestartKillSignal() *string

	GetSendSIGHUP() *string

	GetSendSIGKILL() *string

	GetWatchdogSignal() *string
}

var _ ServiceSection = (*ServiceSectionImpl)(nil)

type ServiceSectionImpl struct {
	Name any `pkl:"name"`

	Type *servicetype.ServiceType `pkl:"type"`

	BusName *string `pkl:"busName"`

	Environment map[string]string `pkl:"environment"`

	EnvironmentFile *string `pkl:"environmentFile"`

	ExecCondition *string `pkl:"execCondition"`

	ExecReload *string `pkl:"execReload"`

	ExecStart *string `pkl:"execStart"`

	ExecStartPost *string `pkl:"execStartPost"`

	ExecStartPre *string `pkl:"execStartPre"`

	ExecStop *string `pkl:"execStop"`

	ExecStopPost *string `pkl:"execStopPost"`

	ExitType *serviceexittype.ServiceExitType `pkl:"exitType"`

	FileDescriptorStoreMax *string `pkl:"fileDescriptorStoreMax"`

	FileDescriptorStorePreserve *string `pkl:"fileDescriptorStorePreserve"`

	GuessMainPID *string `pkl:"guessMainPID"`

	NonBlocking *string `pkl:"nonBlocking"`

	NotifyAccess *string `pkl:"notifyAccess"`

	OOMPolicy *string `pkl:"oOMPolicy"`

	OpenFile *string `pkl:"openFile"`

	PIDFile *string `pkl:"pIDFile"`

	ReloadSignal *string `pkl:"reloadSignal"`

	RemainAfterExit *string `pkl:"remainAfterExit"`

	Restart *string `pkl:"restart"`

	RestartForceExitStatus *string `pkl:"restartForceExitStatus"`

	RestartMaxDelaySec *string `pkl:"restartMaxDelaySec"`

	RestartMode *string `pkl:"restartMode"`

	RestartPreventExitStatus *string `pkl:"restartPreventExitStatus"`

	RestartSec *string `pkl:"restartSec"`

	RestartSteps *string `pkl:"restartSteps"`

	RootDirectoryStartOnly *string `pkl:"rootDirectoryStartOnly"`

	RuntimeMaxSec *string `pkl:"runtimeMaxSec"`

	RuntimeRandomizedExtraSec *string `pkl:"runtimeRandomizedExtraSec"`

	Sockets *string `pkl:"sockets"`

	SuccessExitStatus *string `pkl:"successExitStatus"`

	TimeoutAbortSec *uint `pkl:"timeoutAbortSec"`

	TimeoutSec *uint `pkl:"timeoutSec"`

	TimeoutStartFailureMode *uint `pkl:"timeoutStartFailureMode"`

	TimeoutStartSec *uint `pkl:"timeoutStartSec"`

	TimeoutStopFailureMode *uint `pkl:"timeoutStopFailureMode"`

	TimeoutStopSec *uint `pkl:"timeoutStopSec"`

	USBFunctionDescriptors *string `pkl:"uSBFunctionDescriptors"`

	USBFunctionStrings *string `pkl:"uSBFunctionStrings"`

	WatchdogSec *string `pkl:"watchdogSec"`

	FinalKillSignal *string `pkl:"finalKillSignal"`

	KillMode *string `pkl:"killMode"`

	KillSignal *string `pkl:"killSignal"`

	RestartKillSignal *string `pkl:"restartKillSignal"`

	SendSIGHUP *string `pkl:"sendSIGHUP"`

	SendSIGKILL *string `pkl:"sendSIGKILL"`

	WatchdogSignal *string `pkl:"watchdogSignal"`
}

func (rcv *ServiceSectionImpl) GetName() any {
	return rcv.Name
}

func (rcv *ServiceSectionImpl) GetType() *servicetype.ServiceType {
	return rcv.Type
}

func (rcv *ServiceSectionImpl) GetBusName() *string {
	return rcv.BusName
}

func (rcv *ServiceSectionImpl) GetEnvironment() map[string]string {
	return rcv.Environment
}

func (rcv *ServiceSectionImpl) GetEnvironmentFile() *string {
	return rcv.EnvironmentFile
}

func (rcv *ServiceSectionImpl) GetExecCondition() *string {
	return rcv.ExecCondition
}

func (rcv *ServiceSectionImpl) GetExecReload() *string {
	return rcv.ExecReload
}

func (rcv *ServiceSectionImpl) GetExecStart() *string {
	return rcv.ExecStart
}

func (rcv *ServiceSectionImpl) GetExecStartPost() *string {
	return rcv.ExecStartPost
}

func (rcv *ServiceSectionImpl) GetExecStartPre() *string {
	return rcv.ExecStartPre
}

func (rcv *ServiceSectionImpl) GetExecStop() *string {
	return rcv.ExecStop
}

func (rcv *ServiceSectionImpl) GetExecStopPost() *string {
	return rcv.ExecStopPost
}

func (rcv *ServiceSectionImpl) GetExitType() *serviceexittype.ServiceExitType {
	return rcv.ExitType
}

func (rcv *ServiceSectionImpl) GetFileDescriptorStoreMax() *string {
	return rcv.FileDescriptorStoreMax
}

func (rcv *ServiceSectionImpl) GetFileDescriptorStorePreserve() *string {
	return rcv.FileDescriptorStorePreserve
}

func (rcv *ServiceSectionImpl) GetGuessMainPID() *string {
	return rcv.GuessMainPID
}

func (rcv *ServiceSectionImpl) GetNonBlocking() *string {
	return rcv.NonBlocking
}

func (rcv *ServiceSectionImpl) GetNotifyAccess() *string {
	return rcv.NotifyAccess
}

func (rcv *ServiceSectionImpl) GetOOMPolicy() *string {
	return rcv.OOMPolicy
}

func (rcv *ServiceSectionImpl) GetOpenFile() *string {
	return rcv.OpenFile
}

func (rcv *ServiceSectionImpl) GetPIDFile() *string {
	return rcv.PIDFile
}

func (rcv *ServiceSectionImpl) GetReloadSignal() *string {
	return rcv.ReloadSignal
}

func (rcv *ServiceSectionImpl) GetRemainAfterExit() *string {
	return rcv.RemainAfterExit
}

func (rcv *ServiceSectionImpl) GetRestart() *string {
	return rcv.Restart
}

func (rcv *ServiceSectionImpl) GetRestartForceExitStatus() *string {
	return rcv.RestartForceExitStatus
}

func (rcv *ServiceSectionImpl) GetRestartMaxDelaySec() *string {
	return rcv.RestartMaxDelaySec
}

func (rcv *ServiceSectionImpl) GetRestartMode() *string {
	return rcv.RestartMode
}

func (rcv *ServiceSectionImpl) GetRestartPreventExitStatus() *string {
	return rcv.RestartPreventExitStatus
}

func (rcv *ServiceSectionImpl) GetRestartSec() *string {
	return rcv.RestartSec
}

func (rcv *ServiceSectionImpl) GetRestartSteps() *string {
	return rcv.RestartSteps
}

func (rcv *ServiceSectionImpl) GetRootDirectoryStartOnly() *string {
	return rcv.RootDirectoryStartOnly
}

func (rcv *ServiceSectionImpl) GetRuntimeMaxSec() *string {
	return rcv.RuntimeMaxSec
}

func (rcv *ServiceSectionImpl) GetRuntimeRandomizedExtraSec() *string {
	return rcv.RuntimeRandomizedExtraSec
}

func (rcv *ServiceSectionImpl) GetSockets() *string {
	return rcv.Sockets
}

func (rcv *ServiceSectionImpl) GetSuccessExitStatus() *string {
	return rcv.SuccessExitStatus
}

func (rcv *ServiceSectionImpl) GetTimeoutAbortSec() *uint {
	return rcv.TimeoutAbortSec
}

func (rcv *ServiceSectionImpl) GetTimeoutSec() *uint {
	return rcv.TimeoutSec
}

func (rcv *ServiceSectionImpl) GetTimeoutStartFailureMode() *uint {
	return rcv.TimeoutStartFailureMode
}

func (rcv *ServiceSectionImpl) GetTimeoutStartSec() *uint {
	return rcv.TimeoutStartSec
}

func (rcv *ServiceSectionImpl) GetTimeoutStopFailureMode() *uint {
	return rcv.TimeoutStopFailureMode
}

func (rcv *ServiceSectionImpl) GetTimeoutStopSec() *uint {
	return rcv.TimeoutStopSec
}

func (rcv *ServiceSectionImpl) GetUSBFunctionDescriptors() *string {
	return rcv.USBFunctionDescriptors
}

func (rcv *ServiceSectionImpl) GetUSBFunctionStrings() *string {
	return rcv.USBFunctionStrings
}

func (rcv *ServiceSectionImpl) GetWatchdogSec() *string {
	return rcv.WatchdogSec
}

func (rcv *ServiceSectionImpl) GetFinalKillSignal() *string {
	return rcv.FinalKillSignal
}

func (rcv *ServiceSectionImpl) GetKillMode() *string {
	return rcv.KillMode
}

func (rcv *ServiceSectionImpl) GetKillSignal() *string {
	return rcv.KillSignal
}

func (rcv *ServiceSectionImpl) GetRestartKillSignal() *string {
	return rcv.RestartKillSignal
}

func (rcv *ServiceSectionImpl) GetSendSIGHUP() *string {
	return rcv.SendSIGHUP
}

func (rcv *ServiceSectionImpl) GetSendSIGKILL() *string {
	return rcv.SendSIGKILL
}

func (rcv *ServiceSectionImpl) GetWatchdogSignal() *string {
	return rcv.WatchdogSignal
}
