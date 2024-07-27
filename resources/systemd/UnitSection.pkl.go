// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

type UnitSection interface {
	Section

	GetName() any

	GetAfter() *string

	GetAllowIsolate() *string

	GetBefore() *string

	GetBindsTo() *string

	GetCollectMode() *string

	GetConflicts() *string

	GetDefaultDependencies() *string

	GetDescription() *string

	GetDocumentation() *string

	GetFailureAction() *string

	GetFailureActionExitStatus() *string

	GetIgnoreOnIsolate() *string

	GetJobRunningTimeoutSec() *string

	GetJobTimeoutAction() *string

	GetJobTimeoutRebootArgument() *string

	GetJobTimeoutSec() *string

	GetJoinsNamespaceOf() *string

	GetOnFailure() *string

	GetOnFailureJobMode() *string

	GetOnSuccess() *string

	GetOnSuccessJobMode() *string

	GetPartOf() *string

	GetPropagatesReloadTo() *string

	GetPropagatesStopTo() *string

	GetRebootArgument() *string

	GetRefuseManualStart() *string

	GetRefuseManualStop() *string

	GetReloadPropagatedFrom() *string

	GetRequires() *string

	GetRequiresMountsFor() *string

	GetRequisite() *string

	GetSourcePath() *string

	GetStartLimitAction() *string

	GetStartLimitBurst() *string

	GetStartLimitIntervalSec() *string

	GetStopPropagatedFrom() *string

	GetStopWhenUnneeded() *string

	GetSuccessAction() *string

	GetSuccessActionExitStatus() *string

	GetSurviveFinalKillSignal() *string

	GetUpholds() *string

	GetWants() *string

	GetWantsMountsFor() *string
}

var _ UnitSection = (*UnitSectionImpl)(nil)

type UnitSectionImpl struct {
	Name any `pkl:"name"`

	After *string `pkl:"after"`

	AllowIsolate *string `pkl:"allowIsolate"`

	Before *string `pkl:"before"`

	BindsTo *string `pkl:"bindsTo"`

	CollectMode *string `pkl:"collectMode"`

	Conflicts *string `pkl:"conflicts"`

	DefaultDependencies *string `pkl:"defaultDependencies"`

	Description *string `pkl:"description"`

	Documentation *string `pkl:"documentation"`

	FailureAction *string `pkl:"failureAction"`

	FailureActionExitStatus *string `pkl:"failureActionExitStatus"`

	IgnoreOnIsolate *string `pkl:"ignoreOnIsolate"`

	JobRunningTimeoutSec *string `pkl:"jobRunningTimeoutSec"`

	JobTimeoutAction *string `pkl:"jobTimeoutAction"`

	JobTimeoutRebootArgument *string `pkl:"jobTimeoutRebootArgument"`

	JobTimeoutSec *string `pkl:"jobTimeoutSec"`

	JoinsNamespaceOf *string `pkl:"joinsNamespaceOf"`

	OnFailure *string `pkl:"onFailure"`

	OnFailureJobMode *string `pkl:"onFailureJobMode"`

	OnSuccess *string `pkl:"onSuccess"`

	OnSuccessJobMode *string `pkl:"onSuccessJobMode"`

	PartOf *string `pkl:"partOf"`

	PropagatesReloadTo *string `pkl:"propagatesReloadTo"`

	PropagatesStopTo *string `pkl:"propagatesStopTo"`

	RebootArgument *string `pkl:"rebootArgument"`

	RefuseManualStart *string `pkl:"refuseManualStart"`

	RefuseManualStop *string `pkl:"refuseManualStop"`

	ReloadPropagatedFrom *string `pkl:"reloadPropagatedFrom"`

	Requires *string `pkl:"requires"`

	RequiresMountsFor *string `pkl:"requiresMountsFor"`

	Requisite *string `pkl:"requisite"`

	SourcePath *string `pkl:"sourcePath"`

	StartLimitAction *string `pkl:"startLimitAction"`

	StartLimitBurst *string `pkl:"startLimitBurst"`

	StartLimitIntervalSec *string `pkl:"startLimitIntervalSec"`

	StopPropagatedFrom *string `pkl:"stopPropagatedFrom"`

	StopWhenUnneeded *string `pkl:"stopWhenUnneeded"`

	SuccessAction *string `pkl:"successAction"`

	SuccessActionExitStatus *string `pkl:"successActionExitStatus"`

	SurviveFinalKillSignal *string `pkl:"surviveFinalKillSignal"`

	Upholds *string `pkl:"upholds"`

	Wants *string `pkl:"wants"`

	WantsMountsFor *string `pkl:"wantsMountsFor"`
}

func (rcv *UnitSectionImpl) GetName() any {
	return rcv.Name
}

func (rcv *UnitSectionImpl) GetAfter() *string {
	return rcv.After
}

func (rcv *UnitSectionImpl) GetAllowIsolate() *string {
	return rcv.AllowIsolate
}

func (rcv *UnitSectionImpl) GetBefore() *string {
	return rcv.Before
}

func (rcv *UnitSectionImpl) GetBindsTo() *string {
	return rcv.BindsTo
}

func (rcv *UnitSectionImpl) GetCollectMode() *string {
	return rcv.CollectMode
}

func (rcv *UnitSectionImpl) GetConflicts() *string {
	return rcv.Conflicts
}

func (rcv *UnitSectionImpl) GetDefaultDependencies() *string {
	return rcv.DefaultDependencies
}

func (rcv *UnitSectionImpl) GetDescription() *string {
	return rcv.Description
}

func (rcv *UnitSectionImpl) GetDocumentation() *string {
	return rcv.Documentation
}

func (rcv *UnitSectionImpl) GetFailureAction() *string {
	return rcv.FailureAction
}

func (rcv *UnitSectionImpl) GetFailureActionExitStatus() *string {
	return rcv.FailureActionExitStatus
}

func (rcv *UnitSectionImpl) GetIgnoreOnIsolate() *string {
	return rcv.IgnoreOnIsolate
}

func (rcv *UnitSectionImpl) GetJobRunningTimeoutSec() *string {
	return rcv.JobRunningTimeoutSec
}

func (rcv *UnitSectionImpl) GetJobTimeoutAction() *string {
	return rcv.JobTimeoutAction
}

func (rcv *UnitSectionImpl) GetJobTimeoutRebootArgument() *string {
	return rcv.JobTimeoutRebootArgument
}

func (rcv *UnitSectionImpl) GetJobTimeoutSec() *string {
	return rcv.JobTimeoutSec
}

func (rcv *UnitSectionImpl) GetJoinsNamespaceOf() *string {
	return rcv.JoinsNamespaceOf
}

func (rcv *UnitSectionImpl) GetOnFailure() *string {
	return rcv.OnFailure
}

func (rcv *UnitSectionImpl) GetOnFailureJobMode() *string {
	return rcv.OnFailureJobMode
}

func (rcv *UnitSectionImpl) GetOnSuccess() *string {
	return rcv.OnSuccess
}

func (rcv *UnitSectionImpl) GetOnSuccessJobMode() *string {
	return rcv.OnSuccessJobMode
}

func (rcv *UnitSectionImpl) GetPartOf() *string {
	return rcv.PartOf
}

func (rcv *UnitSectionImpl) GetPropagatesReloadTo() *string {
	return rcv.PropagatesReloadTo
}

func (rcv *UnitSectionImpl) GetPropagatesStopTo() *string {
	return rcv.PropagatesStopTo
}

func (rcv *UnitSectionImpl) GetRebootArgument() *string {
	return rcv.RebootArgument
}

func (rcv *UnitSectionImpl) GetRefuseManualStart() *string {
	return rcv.RefuseManualStart
}

func (rcv *UnitSectionImpl) GetRefuseManualStop() *string {
	return rcv.RefuseManualStop
}

func (rcv *UnitSectionImpl) GetReloadPropagatedFrom() *string {
	return rcv.ReloadPropagatedFrom
}

func (rcv *UnitSectionImpl) GetRequires() *string {
	return rcv.Requires
}

func (rcv *UnitSectionImpl) GetRequiresMountsFor() *string {
	return rcv.RequiresMountsFor
}

func (rcv *UnitSectionImpl) GetRequisite() *string {
	return rcv.Requisite
}

func (rcv *UnitSectionImpl) GetSourcePath() *string {
	return rcv.SourcePath
}

func (rcv *UnitSectionImpl) GetStartLimitAction() *string {
	return rcv.StartLimitAction
}

func (rcv *UnitSectionImpl) GetStartLimitBurst() *string {
	return rcv.StartLimitBurst
}

func (rcv *UnitSectionImpl) GetStartLimitIntervalSec() *string {
	return rcv.StartLimitIntervalSec
}

func (rcv *UnitSectionImpl) GetStopPropagatedFrom() *string {
	return rcv.StopPropagatedFrom
}

func (rcv *UnitSectionImpl) GetStopWhenUnneeded() *string {
	return rcv.StopWhenUnneeded
}

func (rcv *UnitSectionImpl) GetSuccessAction() *string {
	return rcv.SuccessAction
}

func (rcv *UnitSectionImpl) GetSuccessActionExitStatus() *string {
	return rcv.SuccessActionExitStatus
}

func (rcv *UnitSectionImpl) GetSurviveFinalKillSignal() *string {
	return rcv.SurviveFinalKillSignal
}

func (rcv *UnitSectionImpl) GetUpholds() *string {
	return rcv.Upholds
}

func (rcv *UnitSectionImpl) GetWants() *string {
	return rcv.Wants
}

func (rcv *UnitSectionImpl) GetWantsMountsFor() *string {
	return rcv.WantsMountsFor
}
