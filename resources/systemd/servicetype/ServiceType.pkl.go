// Code generated from Pkl module `mikea.declix.resources.systemd`. DO NOT EDIT.
package servicetype

import (
	"encoding"
	"fmt"
)

type ServiceType string

const (
	Simple       ServiceType = "simple"
	Exec         ServiceType = "exec"
	Forking      ServiceType = "forking"
	Oneshot      ServiceType = "oneshot"
	Dbus         ServiceType = "dbus"
	Notify       ServiceType = "notify"
	NotifyReload ServiceType = "notify-reload"
	Idle         ServiceType = "idle"
)

// String returns the string representation of ServiceType
func (rcv ServiceType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ServiceType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ServiceType.
func (rcv *ServiceType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "simple":
		*rcv = Simple
	case "exec":
		*rcv = Exec
	case "forking":
		*rcv = Forking
	case "oneshot":
		*rcv = Oneshot
	case "dbus":
		*rcv = Dbus
	case "notify":
		*rcv = Notify
	case "notify-reload":
		*rcv = NotifyReload
	case "idle":
		*rcv = Idle
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ServiceType`, str)
	}
	return nil
}
