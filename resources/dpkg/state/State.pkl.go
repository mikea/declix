// Code generated from Pkl module `mikea.declix.resources.dpkg`. DO NOT EDIT.
package state

import (
	"encoding"
	"fmt"
)

type State string

const (
	Installed State = "installed"
	Missing   State = "missing"
)

// String returns the string representation of State
func (rcv State) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(State)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for State.
func (rcv *State) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "installed":
		*rcv = Installed
	case "missing":
		*rcv = Missing
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid State`, str)
	}
	return nil
}
