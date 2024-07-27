// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

type InstallSection interface {
	Section

	GetName() any

	GetAlias() *string

	GetAlso() *string

	GetDefaultInstance() *string

	GetRequiredBy() *string

	GetUpheldBy() *string

	GetWantedBy() *string
}

var _ InstallSection = (*InstallSectionImpl)(nil)

type InstallSectionImpl struct {
	Name any `pkl:"name"`

	Alias *string `pkl:"alias"`

	Also *string `pkl:"also"`

	DefaultInstance *string `pkl:"defaultInstance"`

	RequiredBy *string `pkl:"requiredBy"`

	UpheldBy *string `pkl:"upheldBy"`

	WantedBy *string `pkl:"wantedBy"`
}

func (rcv *InstallSectionImpl) GetName() any {
	return rcv.Name
}

func (rcv *InstallSectionImpl) GetAlias() *string {
	return rcv.Alias
}

func (rcv *InstallSectionImpl) GetAlso() *string {
	return rcv.Also
}

func (rcv *InstallSectionImpl) GetDefaultInstance() *string {
	return rcv.DefaultInstance
}

func (rcv *InstallSectionImpl) GetRequiredBy() *string {
	return rcv.RequiredBy
}

func (rcv *InstallSectionImpl) GetUpheldBy() *string {
	return rcv.UpheldBy
}

func (rcv *InstallSectionImpl) GetWantedBy() *string {
	return rcv.WantedBy
}
