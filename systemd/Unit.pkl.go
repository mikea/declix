// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

import "mikea/declix/resources"

type Unit interface {
	resources.Resource

	GetName() string

	GetUser() *string

	GetSystemctl() string

	GetIsEnabled() string

	GetIsActive() string

	GetStateCmd() string

	GetCmds() *UnitStateScripts
}
