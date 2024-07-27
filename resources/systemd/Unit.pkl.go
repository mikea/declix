// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

import "mikea/declix/resources"

type Unit interface {
	resources.Resource

	GetName() string

	GetSystemctl() string

	GetIsEnabled() string

	GetIsActive() string
}
