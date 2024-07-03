package file

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
)

type status struct {
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString() string {
	return pterm.BgCyan.Sprint("NOT IMPLEMENTED")
}

func New(pkl pkl.File) interfaces.Resource {
	return resource{pkl: pkl}
}

type resource struct {
	pkl pkl.File
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() string {
	return pterm.FgGreen.Sprint("exists")
}

// DetermineStatus implements interfaces.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) interfaces.ResouceStatus {
	return status{}
}

// Id implements interfaces.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetPath())
}

// Pkl implements interfaces.Resource.
func (r resource) Pkl() pkl.Resource {
	return r.pkl
}
