package systemd

import (
	"fmt"
	"mikea/declix/interfaces"

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

type action int

// StyledString implements interfaces.Action.
func (a action) StyledString(resource interfaces.Resource) string {
	switch a {
	case toDisable:
		return pterm.FgRed.Sprint("\u2a2f", resource.GetId())
	case toEnable:
		return pterm.FgGreen.Sprint("\u2713", resource.GetId())
	default:
		panic(fmt.Sprintf("unexpected systemd.action: %#v", a))
	}
}

const (
	toEnable action = iota
	toDisable
)

func (s *ServiceImpl) DetermineAction(c interfaces.State, e interfaces.State) (interfaces.Action, error) {
	current := c.(*ServiceState)
	expected := e.(*ServiceState)

	if expected.Enabled {
		if current.Enabled {
			return nil, nil
		}
		return toEnable, nil
	}

	if !current.Enabled {
		return nil, nil
	}
	return toDisable, nil
}

func (s *ServiceImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, c interfaces.State, e interfaces.State) error {
	switch a.(action) {
	case toEnable:
		return executor.Execute(s.Cmds.Enable)
	case toDisable:
		return executor.Execute(s.Cmds.Enable)
	}

	panic(fmt.Sprintf("unexpected systemd.action: %v %+v %+v", a, c, e))
}

func (s *ServiceState) GetStyledString() string {
	if s.Enabled {
		return pterm.FgGreen.Sprint("enabled")
	} else {
		return pterm.FgRed.Sprint("disabled")
	}
}
