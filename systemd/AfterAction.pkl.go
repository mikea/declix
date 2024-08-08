// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package systemd

type AfterAction struct {
	DaemonReload bool `pkl:"daemonReload"`

	Restart *string `pkl:"restart"`

	Reload *string `pkl:"reload"`

	ReloadOrRestart *string `pkl:"reloadOrRestart"`
}
