package dpkg

import (
	"fmt"
	"mikea/declix/content"
	"mikea/declix/interfaces"
	"mikea/declix/resources/dpkg/state"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

type State struct {
	Present bool
	Version string
}

type Action int

const (
	ToInstall Action = iota
	ToRemove  Action = iota
)

// StyledString implements interfaces.Action.
func (a Action) StyledString(resource interfaces.Resource) string {
	switch a {
	case ToInstall:
		return pterm.FgGreen.Sprint("+", resource.GetId())
	case ToRemove:
		return pterm.FgRed.Sprint("-", resource.GetId())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

// DetermineAction implements interfaces.Resource.
func (p PackageImpl) DetermineAction(s interfaces.State, es interfaces.State) (interfaces.Action, error) {
	state := s.(State)
	expectedState := es.(State)

	return DetermineAction(state, expectedState)
}

func DetermineAction(state State, expectedState State) (interfaces.Action, error) {
	if expectedState.Present {
		if state.Present {
			return nil, nil
		}
		return ToInstall, nil
	}

	if state.Present {
		return ToRemove, nil
	}
	return nil, nil
}

// DetermineState implements interfaces.Resource.
func (p PackageImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	return DeterminePackageState(executor, p.GetName())
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (p PackageImpl) ExpectedState() (interfaces.State, error) {
	return State{
		Present: p.State == state.Installed,
		Version: "",
	}, nil
}

// RunAction implements interfaces.Resource.
func (p PackageImpl) RunAction(executor interfaces.CommandExecutor, action interfaces.Action, s interfaces.State, es interfaces.State) error {
	switch action.(Action) {
	case ToInstall:
		io, size, err := content.Open(p.Content)
		if err != nil {
			return err
		}
		defer io.Close()

		tmp, err := executor.UploadTemp(io, size)
		if err != nil {
			return err
		}
		out, err := executor.Run(fmt.Sprintf("sudo dpkg -i %s", tmp))
		if err != nil {
			return fmt.Errorf("error installing package: %w\n%s", err, out)
		}
		out, err = executor.Run(fmt.Sprintf("sudo rm -f %s", tmp))
		if err != nil {
			return fmt.Errorf("error cleaning up: %w\n%s", err, out)
		}
		return nil
	case ToRemove:
		out, err := executor.Run(fmt.Sprintf("sudo dpkg -r %s", p.Name))
		if err != nil {
			return fmt.Errorf("error installing package: %w\n%s", err, out)
		}
		return nil
	}
	panic(fmt.Sprintf("unexpected action: %#v", action))
}

type stateOutput struct {
	Abbrev  string
	Version string
}

func DeterminePackageState(executor interfaces.CommandExecutor, name string) (interfaces.State, error) {
	cmdOut, err := executor.Run(fmt.Sprintf(
		"dpkg-query -W -f='abbrev: ${db:Status-Abbrev}\nversion: ${Version}\n' %s || { [[ $? -eq 1 ]] && echo 'abbrev: uu'; }",
		name))

	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return nil, err
		}

		if e.ExitStatus() == 1 {
			return State{
				Present: false,
				Version: "",
			}, nil
		} else {
			return nil, e
		}
	}

	output := stateOutput{}
	if err := yaml.Unmarshal([]byte(cmdOut), &output); err != nil {
		return nil, err
	}

	return State{
		Present: output.Present(),
		Version: output.Version,
	}, nil
}

// StyledString implements interfaces.ResouceStatus.
func (s State) StyledString(resource interfaces.Resource) string {
	if s.Present {
		if s.Version == "" {
			return pterm.FgGreen.Sprint("installed")
		}
		return pterm.FgGreen.Sprint(s.Version)
	}
	return pterm.FgRed.Sprint("missing")
}

func (s stateOutput) Present() bool {
	if len(s.Abbrev) >= 2 && s.Abbrev[1] == 'i' {
		return true
	}
	return false
}
