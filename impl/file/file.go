package file

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func New(pkl pkl.File) interfaces.Resource {
	return resource{pkl: pkl}
}

type resource struct {
	pkl pkl.File
}

// RunAction implements interfaces.Resource.
func (r resource) RunAction(executor interfaces.CommandExcutor, actio interfaces.Action, status interfaces.Status) error {
	tmp, err := executor.Run("mktemp")
	if err != nil {
		return err
	}
	tmp = strings.TrimSuffix(tmp, "\n")
	defer executor.Run(fmt.Sprintf("rm -f %s", tmp))

	content, err := os.Open(r.pkl.GetContentFile())
	if err != nil {
		return err
	}
	defer content.Close()

	err = executor.Upload(*content, tmp, "0655")
	if err != nil {
		return fmt.Errorf("error uploading file: %w", err)
	}

	_, err = executor.Run(fmt.Sprintf("sudo -S cp %s %s", tmp, r.pkl.GetPath()))
	if err != nil {
		return fmt.Errorf("error copyinh file: %w", err)
	}

	return nil
}

// DetermineAction implements interfaces.Resource.
func (r resource) DetermineAction(executor interfaces.CommandExcutor, s interfaces.Status) (interfaces.Action, error) {
	status := s.(status)

	if status.Exists {
		return nil, nil
	}

	return ToCreate, nil
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
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) (interfaces.Status, error) {
	out, err := executor.Run(fmt.Sprintf("if [ ! -f \"%s\" ]; then echo \"exists: false\"; else echo \"exists: true\"; fi", r.pkl.GetPath()))
	if err != nil {
		return nil, err
	}

	status := status{}
	yaml.Unmarshal([]byte(out), &status)
	return status, nil
}

// Id implements interfaces.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetPath())
}

// Pkl implements interfaces.Resource.
func (r resource) Pkl() pkl.Resource {
	return r.pkl
}
