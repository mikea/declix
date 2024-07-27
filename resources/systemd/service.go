package systemd

import (
	"mikea/declix/interfaces"
	"strings"

	"github.com/pterm/pterm"
)

func (s *ServiceImpl) ExpectedState() (interfaces.State, error) { return s.State, nil }

func (s *ServiceImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	state := &ServiceState{}
	if err := executor.Evaluate(s.StateCmd, state); err != nil {
		return nil, err
	}
	return state, nil
}

// StyledString implements interfaces.Action.
func (a *action) StyledString(resource interfaces.Resource) string {
	var result = ""

	if a.enable != nil {
		if *a.enable {
			result = result + pterm.FgGreen.Sprint("+")
		} else {
			result = result + pterm.FgRed.Sprint("-")
		}
	}

	if a.start != nil {
		if *a.start {
			result = result + pterm.FgGreen.Sprint(">")
		} else {
			result = result + pterm.FgRed.Sprint("_")
		}
	}

	return result + resource.GetId()
}

type action struct {
	enable *bool
	start  *bool
}

func (s *ServiceImpl) DetermineAction(c interfaces.State, e interfaces.State) (interfaces.Action, error) {
	current := c.(*ServiceState)
	expected := e.(*ServiceState)

	a := &action{}

	if expected.Enabled != nil && *expected.Enabled != *current.Enabled {
		a.enable = expected.Enabled
	}
	if expected.Active != nil && *expected.Active != *current.Active {
		a.start = expected.Active
	}

	return a, nil
}

func (s *ServiceImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, c interfaces.State, e interfaces.State) error {
	action := a.(*action)

	if action.enable != nil {
		if err := executor.Execute(*s.Cmds.Enable); err != nil {
			return err
		}
	}
	if action.start != nil {
		if err := executor.Execute(*s.Cmds.Start); err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceState) GetStyledString() string {
	var result []string

	if s.Enabled != nil {
		if *s.Enabled {
			result = append(result, pterm.FgGreen.Sprint("enabled"))
		} else {
			result = append(result, pterm.FgRed.Sprint("disabled"))
		}
	}

	if s.Active != nil {
		if *s.Active {
			result = append(result, pterm.FgGreen.Sprint("active"))
		} else {
			result = append(result, pterm.FgRed.Sprint("inactive"))
		}
	}

	return strings.Join(result, " ")
}
