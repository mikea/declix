package users

import (
	"fmt"
	"mikea/declix/interfaces"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func (g *GroupImpl) ExpectedState() (interfaces.State, error) { return g.State.(interfaces.State), nil }

type groupStateOutput struct {
	Present bool
	Gid     int
}

func (g *GroupImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	out, err := executor.Run(fmt.Sprintf(`
		name="%s"
		if  entry=$(getent group "$name"); then
		    IFS=':' read -r name passwd gid members <<< "$entry"
			echo "present: true"
			echo "gid: $gid"
		else 
			echo "present: false"
		fi`,
		g.Name,
	))
	if err != nil {
		return nil, err
	}
	output := groupStateOutput{}
	if err := yaml.Unmarshal([]byte(out), &output); err != nil {
		return nil, err
	}

	if !output.Present {
		return &GroupMissing{}, nil
	}

	return &GroupPresent{
		Gid: output.Gid,
	}, err
}

type groupAction int

const (
	groupCreate groupAction = iota
	groupUpdate
	groupDelete
)

// StyledString implements interfaces.Action.
func (a groupAction) StyledString(resource interfaces.Resource) string {
	switch a {
	case groupCreate:
		return pterm.FgGreen.Sprint("+", resource.GetId())
	case groupDelete:
		return pterm.FgRed.Sprint("-", resource.GetId())
	case groupUpdate:
		return pterm.FgYellow.Sprint("~", resource.GetId())
	default:
		panic(fmt.Sprintf("unexpected group action: %#v", a))
	}
}

func (g *GroupImpl) DetermineAction(s interfaces.State, e interfaces.State) (interfaces.Action, error) {
	switch expectedState := e.(type) {
	case *GroupMissing:
		if _, ok := s.(*GroupMissing); ok {
			return nil, nil
		}
		return groupDelete, nil
	case *GroupPresent:
		if state, ok := s.(*GroupPresent); ok {
			if expectedState.Gid != state.Gid {
				return groupUpdate, nil
			}
			return nil, nil
		}
		return groupCreate, nil
	}
	panic(fmt.Sprintf("wrong state %T", e))
}

func (g *GroupImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, e interfaces.State) error {
	action := a.(groupAction)

	switch action {
	case groupCreate:
		expected := e.(*GroupPresent)
		out, err := executor.Run(fmt.Sprintf("sudo groupadd -g %d %s", expected.Gid, g.Name))
		if err != nil {
			return fmt.Errorf("error creating group: %w\n%s", err, out)
		}
		return nil
	case groupDelete:
		out, err := executor.Run(fmt.Sprintf("sudo groupdel %s", g.Name))
		if err != nil {
			return fmt.Errorf("error deleting group: %w\n%s", err, out)
		}
		return nil
	case groupUpdate:
		expected := e.(*GroupPresent)
		out, err := executor.Run(fmt.Sprintf("sudo groupmod -g %d %s", expected.Gid, g.Name))
		if err != nil {
			return fmt.Errorf("error updating group: %w\n%s", err, out)
		}
		return nil
	default:
		panic(fmt.Sprintf("unexpected group action: %#v", action))
	}
}

func (state *GroupPresent) StyledString(r interfaces.Resource) string {
	return pterm.FgGreen.Sprintf("%s %d", r.(*GroupImpl).Name, state.Gid)
}

func (state *GroupMissing) StyledString(r interfaces.Resource) string {
	return pterm.FgRed.Sprint("missing")
}
