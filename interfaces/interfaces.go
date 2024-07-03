package interfaces

import (
	"mikea/declix/pkl"
)

type Resource interface {
	Id() string
	Pkl() pkl.Resource

	ExpectedStatusStyledString() string
	DetermineStatus(executor CommandExcutor) ResouceStatus
}

type ResouceStatus interface {
	StyledString() string
}

type CommandExcutor interface {
	Run(command string) (string, error)
	Close() error
}
