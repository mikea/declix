package interfaces

import (
	"mikea/declix/pkl"
)

type Resource interface {
	Id() string
	Pkl() pkl.Resource

	ExpectedStatusStyledString() string
	DetermineStatus(executor CommandExcutor) Status
	DetermineAction(executor CommandExcutor) Action
}

type Status interface {
	StyledString(resource Resource) string
}

type CommandExcutor interface {
	Run(command string) (string, error)
	Close() error
}

type Action interface {
	StyledString(resource Resource) string
}
