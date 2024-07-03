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

// DetermineAction implements interfaces.Resource.
func (r resource) DetermineAction(executor interfaces.CommandExcutor) interfaces.Action {
	status := r.DetermineStatus(executor).(status)

	if status.Exists {
		return nil
	}

	return ToCreate
}

type status struct {
	Exists bool
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString(resource interfaces.Resource) string {
	if !s.Exists {
		return pterm.FgRed.Sprint("missing")
	} else {
		return pterm.FgGreen.Sprint("exists")
	}
}

type action int

// StyledString implements interfaces.Action.
func (a action) StyledString(resource interfaces.Resource) string {
	switch a {
	case ToCreate:
		return pterm.FgGreen.Sprint("+", resource.Id())
		// case ToRemove:
		// 	return pterm.FgRed.Sprint("-", resource.Id())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

const (
	ToCreate action = iota
)

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() string {
	return pterm.FgGreen.Sprint("exists")
}

// DetermineStatus implements interfaces.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) interfaces.Status {
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
