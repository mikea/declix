package interfaces

import (
	"io"
)

type Resource interface {
	GetId() string

	ExpectedState() (State, error)
	DetermineState(executor CommandExecutor) (State, error)
	DetermineAction(state State, expectedState State) (Action, error)
	RunAction(executor CommandExecutor, action Action, state State, expectedState State) error
}

type State interface {
	GetStyledString() string
}

type CommandExecutor interface {
	Close() error
	MkTemp() (string, error)
	Run(command string) (string, error)

	Execute(command string) error
	Evaluate(command string, out any) error

	Upload(content io.Reader, remotePath string, permissions string, size int64) error
	UploadTemp(content io.Reader, size int64) (string, error)
}

type Action interface {
	StyledString(resource Resource) string
}
