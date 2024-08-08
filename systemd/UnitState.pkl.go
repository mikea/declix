// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type UnitState interface {
	GetEnabled() *bool

	GetActive() *bool
}

var _ UnitState = (*UnitStateImpl)(nil)

type UnitStateImpl struct {
	Enabled *bool `pkl:"enabled"`

	Active *bool `pkl:"active"`
}

func (rcv *UnitStateImpl) GetEnabled() *bool {
	return rcv.Enabled
}

func (rcv *UnitStateImpl) GetActive() *bool {
	return rcv.Active
}
