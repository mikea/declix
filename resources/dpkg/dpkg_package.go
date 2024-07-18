package dpkg

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

type Status struct {
	Abbrev  string
	Version string
}

type state int

const (
	Missing state = iota
	Installed
)

// DetermineAction implements interfaces.Resource.
func (p PackageImpl) DetermineAction(executor interfaces.CommandExcutor, status interfaces.Status) (interfaces.Action, error) {
	panic("unimplemented")
}

// DetermineStatus implements interfaces.Resource.
func (p PackageImpl) DetermineStatus(executor interfaces.CommandExcutor) (interfaces.Status, error) {
	return DeterminePackageStatus(executor, p.GetName())
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (p PackageImpl) ExpectedStatusStyledString() (string, error) {
	panic("unimplemented")
}

// Id implements interfaces.Resource.
func (p PackageImpl) Id() string {
	return fmt.Sprintf("%s:%s", p.Type, p.GetName())
}

// Pkl implements interfaces.Resource.
func (p PackageImpl) Pkl() resources.Resource {
	panic("unimplemented")
}

// RunAction implements interfaces.Resource.
func (p PackageImpl) RunAction(executor interfaces.CommandExcutor, action interfaces.Action, status interfaces.Status) error {
	panic("unimplemented")
}

func DeterminePackageStatus(executor interfaces.CommandExcutor, name string) (interfaces.Status, error) {
	out, err := executor.Run(fmt.Sprintf(
		"dpkg-query -W -f='abbrev: ${db:Status-Abbrev}\nversion: ${Version}\n' %s || { [[ $? -eq 1 ]] && echo 'abbrev: uu'; }",
		name))

	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return nil, err
		}

		if e.ExitStatus() == 1 {
			return Status{
				Abbrev: "uu",
			}, nil
		} else {
			return nil, e
		}
	}

	status := Status{}
	yaml.Unmarshal([]byte(out), &status)
	return status, nil
}

// StyledString implements interfaces.ResouceStatus.
func (s Status) StyledString(resource interfaces.Resource) string {
	switch s.PackageState() {
	case Installed:
		return pterm.FgGreen.Sprint(s.Version)
	case Missing:
		return pterm.FgRed.Sprint("missing")
	}
	panic(fmt.Sprintf("unexpected apt_package.state: %#v", s.Abbrev))
}

func (s Status) PackageState() state {
	if len(s.Abbrev) >= 2 && s.Abbrev[1] == 'i' {
		return Installed
	}
	return Missing
}
