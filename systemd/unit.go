package systemd

import (
	"mikea/declix/interfaces"
	"strings"

	"github.com/pterm/pterm"
)

func (a *unitAction) empty() bool {
	return a.enable == nil && a.start == nil
}

func (a *unitAction) GetStyledString(resource interfaces.Resource) string {
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

type unitAction struct {
	enable *bool
	start  *bool
}

func (action *unitAction) Run(executor interfaces.CommandExecutor, r interfaces.Resource, c interfaces.State, e interfaces.State) error {
	u := r.(Unit)

	if action.enable != nil {
		if err := executor.Execute(*u.GetCmds().Enable); err != nil {
			return err
		}
	}
	if action.start != nil {
		if err := executor.Execute(*u.GetCmds().Start); err != nil {
			return err
		}
	}
	return nil
}

func determineUnitState(u Unit, executor interfaces.CommandExecutor) (interfaces.State, error) {
	state := &UnitStateImpl{}
	if err := executor.Evaluate(u.GetStateCmd(), state); err != nil {
		return nil, err
	}
	return state, nil
}

func determineUnitAction(c interfaces.State, e interfaces.State) (interfaces.Action, error) {
	current := c.(*UnitStateImpl)
	expected := e.(*UnitStateImpl)

	a := &unitAction{}

	if expected.Enabled != nil && *expected.Enabled != *current.Enabled {
		a.enable = expected.Enabled
	}
	if expected.Active != nil && *expected.Active != *current.Active {
		a.start = expected.Active
	}

	if a.empty() {
		return nil, nil
	}

	return a, nil
}

func (s *UnitStateImpl) GetStyledString() string {
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
