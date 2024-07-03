package apt_package

import (
	"fmt"
	"log"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
)

type resource struct {
	pkl pkl.Package
}

type state int

const (
	Error state = iota
	Missing
	Installed
)

type status struct {
	state   state
	version string
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString() string {
	switch s.state {

	case Error:
		return pterm.BgRed.Sprint("ERROR")
	case Installed:
		return pterm.FgGreen.Sprint(s.version)
	case Missing:
		return pterm.FgRed.Sprint("missing")
	}
	panic(fmt.Sprintf("unexpected apt_package.state: %#v", s.state))
}

func New(pkl pkl.Package) interfaces.Resource {
	return resource{pkl: pkl}
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() string {
	switch r.pkl.GetStatus() {
	case "installed":
		return pterm.FgGreen.Sprint("installed")
	case "missing":
		return pterm.FgRed.Sprint("missing")
	}

	panic(fmt.Sprintf("unexpected status: %#v", r.pkl.GetStatus()))
}

// DetermineStatus implements impl.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) interfaces.ResouceStatus {
	output, err := executor.Run(fmt.Sprintf("dpkg-query -W -f='${Version}' %s", r.pkl.GetName()))

	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return status{
				state: Error,
			}
		}

		if e.ExitStatus() == 1 {
			return status{
				state: Missing,
			}
		} else {
			log.Fatalf("Exit error: %v", e)
			return status{
				state: Error,
			}
		}
	}

	if output == "" {
		return status{
			state: Missing,
		}
	}

	return status{
		state:   Installed,
		version: output,
	}
}

// Id implements impl.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetName())
}

// Pkl implements impl.Resource.
func (r resource) Pkl() pkl.Resource {
	return r.pkl
}
