package interfaces

import (
	"io"
	"mikea/declix/pkl"
)

type Resource interface {
	Id() string
	Pkl() pkl.Resource

	ExpectedStatusStyledString() (string, error)
	DetermineStatus(executor CommandExcutor) (Status, error)
	DetermineAction(executor CommandExcutor, status Status) (Action, error)
	RunAction(executor CommandExcutor, action Action, status Status) error
}

type Status interface {
	StyledString(resource Resource) string
}

type CommandExcutor interface {
	Run(command string) (string, error)
	Upload(content io.Reader, remotePath string, permissions string, size int64) error
	Close() error
}

type Action interface {
	StyledString(resource Resource) string
}
