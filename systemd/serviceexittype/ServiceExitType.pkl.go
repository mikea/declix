// Code generated from Pkl module `mikea.declix.systemd`. DO NOT EDIT.
package serviceexittype

import (
	"encoding"
	"fmt"
)

type ServiceExitType string

const (
	Main   ServiceExitType = "main"
	Cgroup ServiceExitType = "cgroup"
)

// String returns the string representation of ServiceExitType
func (rcv ServiceExitType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ServiceExitType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ServiceExitType.
func (rcv *ServiceExitType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "main":
		*rcv = Main
	case "cgroup":
		*rcv = Cgroup
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ServiceExitType`, str)
	}
	return nil
}
