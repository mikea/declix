package file

import (
	"fmt"
	"log"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func New(pkl pkl.File) interfaces.Resource {
	return resource{pkl: pkl}
}

type resource struct {
	pkl pkl.File
}

type status struct {
	Exists bool
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString() string {
	if !s.Exists {
		return pterm.FgRed.Sprint("missing")
	} else {
		return pterm.FgGreen.Sprint("exists")
	}
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() string {
	return pterm.FgGreen.Sprint("exists")
}

// DetermineStatus implements interfaces.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) interfaces.ResouceStatus {
	out, err := executor.Run(fmt.Sprintf("if [ ! -f \"%s\" ]; then echo \"exists: false\"; else echo \"exists: true\"; fi", r.pkl.GetPath()))
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	status := status{}
	yaml.Unmarshal([]byte(out), &status)
	return status
}

// Id implements interfaces.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetPath())
}

// Pkl implements interfaces.Resource.
func (r resource) Pkl() pkl.Resource {
	return r.pkl
}
