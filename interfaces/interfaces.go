package interfaces

import (
	"mikea/declix/pkl"
)

type Resource interface {
	Id() string
	Pkl() pkl.Resource

	ExpectedStatusStyledString() string
	DetermineStatus(executor CommandExcutor) (Status, error)
	DetermineAction(executor CommandExcutor, status Status) (Action, error)
	RunAction(executor CommandExcutor, action Action, status Status) error
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
