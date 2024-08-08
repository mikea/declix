package filesystem

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func (d *DirImpl) ExpectedState() (interfaces.State, error) { return d.State.(interfaces.State), nil }

type groupStateOutput struct {
	Present     bool
	Owner       string
	Group       string
	Permissions string
}

func (d *DirImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	out, err := executor.Run(d.StateCmd)
	if err != nil {
		return nil, err
	}
	output := groupStateOutput{}
	if err := yaml.Unmarshal([]byte(out), &output); err != nil {
		return nil, err
	}

	if !output.Present {
		return &resources.Missing{}, nil
	}

	return &DirPresentImpl{
		Owner:       output.Owner,
		Group:       output.Group,
		Permissions: output.Permissions,
	}, nil
}

func (d *DirImpl) DetermineAction(s interfaces.State, e interfaces.State) (interfaces.Action, error) {
	switch expected := e.(type) {
	case *resources.Missing:
		if _, ok := s.(*resources.Missing); ok {
			return nil, nil
		}
		return ToDelete, nil

	case *DirPresentImpl:
		if state, ok := s.(*DirPresentImpl); ok {
			if expected.Owner != state.Owner || expected.Group != state.Group || expected.Permissions != state.Permissions {
				return ToUpdate, nil
			}
			return nil, nil
		}
		return ToCreate, nil
	}

	panic(fmt.Sprintf("wrong state %T", e))
}

func (d *DirImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, e interfaces.State) error {
	action := a.(action)

	switch action {
	case ToCreate:
		expected := e.(NodePresent)
		out, err := executor.Run(fmt.Sprintf("sudo mkdir \"%s\"", d.GetPath()))
		if err != nil {
			return fmt.Errorf("error creating group: %w\n%s", err, out)
		}
		if err = chown(executor, d.GetPath(), expected.GetOwner()); err != nil {
			return err
		}
		if err = chgrp(executor, d.GetPath(), expected.GetGroup()); err != nil {
			return err
		}
		if err = chmod(executor, d.GetPath(), expected.GetPermissions()); err != nil {
			return err
		}
		return nil
	case ToDelete:
		out, err := executor.Run(fmt.Sprintf("sudo rm -df \"%s\"", d.GetPath()))
		if err != nil {
			return fmt.Errorf("error creating group: %w\n%s", err, out)
		}
		return nil
	case ToUpdate:
		current := s.(NodePresent)
		expected := e.(NodePresent)
		if current.GetOwner() != expected.GetOwner() {
			if err := chown(executor, d.GetPath(), expected.GetOwner()); err != nil {
				return err
			}

		}
		if current.GetGroup() != expected.GetGroup() {
			if err := chgrp(executor, d.GetPath(), expected.GetGroup()); err != nil {
				return err
			}
		}
		if current.GetPermissions() != expected.GetPermissions() {
			if err := chmod(executor, d.GetPath(), expected.GetPermissions()); err != nil {
				return err
			}
		}
		return nil
	default:
		panic(fmt.Sprintf("unexpected filesystem.action: %#v", action))
	}
}

func (state *DirPresentImpl) GetStyledString() string {
	return pterm.FgGreen.Sprintf("%s:%s %s", state.Owner, state.Group, state.Permissions)
}
