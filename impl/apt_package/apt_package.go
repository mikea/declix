package apt_package

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

type resource struct {
	pkl pkl.Package
}

// RunAction implements interfaces.Resource.
func (r resource) RunAction(executor interfaces.CommandExcutor, a interfaces.Action, status interfaces.Status) error {
	switch a.(action) {
	case ToInstall:
		_, err := executor.Run(fmt.Sprintf("sudo -S apt-get install -y --no-upgrade --no-install-recommends %s", r.pkl.GetName()))
		return err
	case ToRemove:
		_, err := executor.Run(fmt.Sprintf("sudo -S apt-get remove -y %s", r.pkl.GetName()))
		return err
	}
	panic("unexpected apt_package.action")
}

// DetermineAction implements interfaces.Resource.
func (r resource) DetermineAction(executor interfaces.CommandExcutor, s interfaces.Status) (interfaces.Action, error) {
	status := s.(status)

	switch r.pkl.GetStatus() {
	case "installed":
		if status.PackageState() == Installed {
			return nil, nil
		}
		return ToInstall, nil
	case "missing":
		if status.PackageState() == Missing {
			return nil, nil
		}
		return ToRemove, nil
	}

	panic(fmt.Sprintf("unexpected status: %#v", r.pkl.GetStatus()))
}

type state int

const (
	Missing state = iota
	Installed
)

type status struct {
	Abbrev  string
	Version string
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString(resource interfaces.Resource) string {
	switch s.PackageState() {
	case Installed:
		return pterm.FgGreen.Sprint(s.Version)
	case Missing:
		return pterm.FgRed.Sprint("missing")
	}
	panic(fmt.Sprintf("unexpected apt_package.state: %#v", s.Abbrev))
}

func (s status) PackageState() state {
	if len(s.Abbrev) >= 2 && s.Abbrev[1] == 'i' {
		return Installed
	}
	return Missing
}

type action int

// StyledString implements interfaces.Action.
func (a action) StyledString(resource interfaces.Resource) string {
	switch a {
	case ToInstall:
		return pterm.FgGreen.Sprint("+", resource.Id())
	case ToRemove:
		return pterm.FgRed.Sprint("-", resource.Id())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

const (
	ToInstall action = iota
	ToRemove  action = iota
)

func New(pkl pkl.Package) interfaces.Resource {
	return resource{pkl: pkl}
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() (string, error) {
	switch r.pkl.GetStatus() {
	case "installed":
		return pterm.FgGreen.Sprint("installed"), nil
	case "missing":
		return pterm.FgRed.Sprint("missing"), nil
	}

	panic(fmt.Sprintf("unexpected status: %#v", r.pkl.GetStatus()))
}

// DetermineStatus implements impl.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) (interfaces.Status, error) {
	out, err := executor.Run(fmt.Sprintf("dpkg-query -W -f='abbrev: ${db:Status-Abbrev}\nversion: ${Version}\n' %s", r.pkl.GetName()))

	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return nil, err
		}

		if e.ExitStatus() == 1 {
			return status{
				Abbrev: "uu",
			}, nil
		} else {
			return nil, e
		}
	}

	status := status{}
	yaml.Unmarshal([]byte(out), &status)
	return status, nil
}

// Id implements impl.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetName())
}

// Pkl implements impl.Resource.
func (r resource) Pkl() pkl.Resource {
	return r.pkl
}
