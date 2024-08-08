package users

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"slices"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func (u *UserImpl) ExpectedState() (interfaces.State, error) { return u.State.(interfaces.State), nil }

type userStateOutput struct {
	Present bool
	Uid     uint16
	Group   string
	Groups  string
}

func (u *UserImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	out, err := executor.Run(fmt.Sprintf(`
		login="%s"
		if id "$login" &>/dev/null; then
			uid=$(id -u "$login")
			group=$(id -gn "$login")
			groups=$(id -Gn "$login")
			echo "present: true"
			echo "uid: $uid"
			echo "group: $group"
			echo "groups: $groups"
		else 
			echo "present: false"
		fi`,
		u.Login,
	))
	if err != nil {
		return nil, err
	}
	output := userStateOutput{}
	if err := yaml.Unmarshal([]byte(out), &output); err != nil {
		return nil, err
	}

	if !output.Present {
		return &resources.Missing{}, nil
	}

	return &UserPresent{
		Uid:    output.Uid,
		Group:  output.Group,
		Groups: strings.Split(output.Groups, " "),
	}, err
}

type userAction int

const (
	userCreate userAction = iota
	userUpdate
	userDelete
)

func (a userAction) GetStyledString(resource interfaces.Resource) string {
	switch a {
	case userCreate:
		return pterm.FgGreen.Sprint("+", resource.GetId())
	case userDelete:
		return pterm.FgRed.Sprint("-", resource.GetId())
	case userUpdate:
		return pterm.FgYellow.Sprint("~", resource.GetId())
	default:
		panic(fmt.Sprintf("unexpected user action: %#v", a))
	}
}

func (u *UserImpl) DetermineAction(s interfaces.State, e interfaces.State) (interfaces.Action, error) {
	switch expectedState := e.(type) {
	case *resources.Missing:
		if _, ok := s.(*resources.Missing); ok {
			return nil, nil
		}
		return userDelete, nil
	case *UserPresent:
		if state, ok := s.(*UserPresent); ok {
			if expectedState.Groups != nil && !slices.Equal(expectedState.Groups, state.Groups) {
				return userUpdate, nil
			}
			if expectedState.Group != state.Group || expectedState.Uid != state.Uid {
				return userUpdate, nil
			}
			return nil, nil
		}
		return userCreate, nil
	}
	panic(fmt.Sprintf("wrong state %T", s))
}

func (u *UserImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, e interfaces.State) error {
	action := a.(userAction)

	switch action {
	case userCreate:
		expected := e.(*UserPresent)
		out, err := executor.Run(fmt.Sprintf("sudo useradd -g %s -u %d -G %s %s",
			expected.Group, expected.Uid, strings.Join(expected.Groups, ","), u.Login))
		if err != nil {
			return fmt.Errorf("error creating user: %w\n%s", err, out)
		}
		return nil
	case userDelete:
		out, err := executor.Run(fmt.Sprintf("sudo userdel %s", u.Login))
		if err != nil {
			return fmt.Errorf("error deleting user: %w\n%s", err, out)
		}
		return nil
	case userUpdate:
		expected := e.(*UserPresent)
		out, err := executor.Run(fmt.Sprintf("sudo usermod -g %s -u %d -G %s %s",
			expected.Group, expected.Uid, strings.Join(expected.Groups, ","), u.Login))
		if err != nil {
			return fmt.Errorf("error updating user: %w\n%s", err, out)
		}
		return nil
	default:
		panic(fmt.Sprintf("unexpected user action: %#v", action))
	}
}

func (state *UserPresent) GetStyledString() string {
	if state.Groups != nil {
		return pterm.FgGreen.Sprintf("%d %s %v", state.Uid, state.Group, state.Groups)
	} else {
		return pterm.FgGreen.Sprintf("%d %s", state.Uid, state.Group)
	}
}
