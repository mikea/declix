// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package systemd

type ServiceFile struct {
	Unit UnitSection `pkl:"unit"`

	Install InstallSection `pkl:"install"`

	Service ServiceSection `pkl:"service"`
}
