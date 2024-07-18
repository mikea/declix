package interfaces

import (
	"io"
	"mikea/declix/resources"
)

type Resource interface {
	Id() string
	Pkl() resources.Resource

	ExpectedState() (State, error)
	DetermineState(executor CommandExcutor) (State, error)
	DetermineAction(executor CommandExcutor, state State, expectedState State) (Action, error)
	RunAction(executor CommandExcutor, action Action, state State, expectedState State) error
}

type State interface {
	StyledString(resource Resource) string
}

type CommandExcutor interface {
	Close() error
	MkTemp() (string, error)
	Run(command string) (string, error)
	Upload(content io.Reader, remotePath string, permissions string, size int64) error
	UploadTemp(content io.Reader, size int64) (string, error)
	UploadTempNoSize(content io.Reader) (string, error)
}

type Action interface {
	StyledString(resource Resource) string
}
