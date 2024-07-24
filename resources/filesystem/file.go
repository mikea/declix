package filesystem

import (
	"fmt"
	"io"
	"mikea/declix/content"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

// RunAction implements interfaces.Resource.
func (f *FileImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, es interfaces.State) error {
	expected := es.(state)

	action := a.(action)
	switch action {
	case ToCreate:
		return f.upload(executor, expected)
	case ToDelete:
		_, err := executor.Run(fmt.Sprintf("sudo rm -f %s", f.Path))
		if err != nil {
			return err
		}
		return nil
	case ToUpdate:
		return f.update(executor, s.(state), expected)
	default:
		panic(fmt.Sprintf("unexpected filesystem.action: %#v", action))
	}
}

func (f *FileImpl) update(executor interfaces.CommandExecutor, current state, expected state) error {
	if current.Sha256 != expected.Sha256 {
		return f.upload(executor, expected)
	}
	if current.Group != expected.Group {
		if err := chgrp(executor, f.Path, expected.Group); err != nil {
			return err
		}
	}
	if current.Owner != expected.Owner {
		if err := chown(executor, f.Path, expected.Owner); err != nil {
			return err
		}
	}
	if current.Permissions != expected.Permissions {
		if err := chmod(executor, f.Path, expected.Permissions); err != nil {
			return err
		}
	}

	return nil
}

func (f *FileImpl) openContent() (io.ReadCloser, int64, error) {
	return content.Open(f.State.(*FilePresentImpl).Content)
}

func (f *FileImpl) upload(executor interfaces.CommandExecutor, expected state) error {
	content, size, err := f.openContent()
	if err != nil {
		return err
	}
	defer content.Close()

	tmp, err := executor.UploadTemp(content, size)
	if err != nil {
		return err
	}

	_, err = executor.Run(fmt.Sprintf("sudo -S mv %s %s", tmp, f.Path))
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	if err := chown(executor, f.Path, expected.Owner); err != nil {
		return err
	}
	if err := chgrp(executor, f.Path, expected.Group); err != nil {
		return err
	}
	if err := chmod(executor, f.Path, expected.Permissions); err != nil {
		return err
	}

	return nil
}

// DetermineAction implements interfaces.Resource.
func (f *FileImpl) DetermineAction(s interfaces.State, es interfaces.State) (interfaces.Action, error) {
	expected := es.(state)
	current := s.(state)

	if expected.Exists {
		if current.Exists {
			if current.Sha256 != expected.Sha256 ||
				current.Owner != expected.Owner ||
				current.Group != expected.Group ||
				current.Permissions != expected.Permissions {
				return ToUpdate, nil
			}

			return nil, nil
		}

		return ToCreate, nil
	}

	if !current.Exists {
		return nil, nil
	}

	return ToDelete, nil
}

// todo: use pkl instead
type state struct {
	Exists      bool
	Sha256      string
	Owner       string
	Group       string
	Permissions string
}

// StyledString implements interfaces.ResouceStatus.
func (s state) StyledString(resource interfaces.Resource) string {
	if !s.Exists {
		return pterm.FgRed.Sprint("missing")
	} else {
		return pterm.FgGreen.Sprint(s.Sha256[:8], " ", s.Owner, ":", s.Group, " ", s.Permissions)
	}
}

type action int

// StyledString implements interfaces.Action.
func (a action) StyledString(resource interfaces.Resource) string {
	switch a {
	case ToCreate:
		return pterm.FgGreen.Sprint("+", resource.GetId())
	case ToUpdate:
		return pterm.FgYellow.Sprint("~", resource.GetId())
	case ToDelete:
		return pterm.FgRed.Sprint("-", resource.GetId())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

const (
	ToCreate action = iota
	ToUpdate
	ToDelete
)

// DetermineState implements interfaces.Resource.
func (f *FileImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	out, err := executor.Run(f.DetermineStateCmd)
	if err != nil {
		return nil, err
	}
	state := state{}
	if err := yaml.Unmarshal([]byte(out), &state); err != nil {
		return nil, err
	}
	return state, nil
}

func (f *FileImpl) ExpectedState() (interfaces.State, error) {
	switch s := f.State.(type) {
	case *resources.Missing:
		return state{
			Exists: false,
		}, nil
	case *FilePresentImpl:
		sha256, err := content.Sha256(f.State.(*FilePresentImpl).Content)
		if err != nil {
			return nil, err
		}

		return state{
			Exists:      true,
			Sha256:      sha256,
			Owner:       s.Owner,
			Group:       s.Group,
			Permissions: s.Permissions,
		}, nil

	default:
		panic(fmt.Sprintf("unsupported state %T", s))

	}

}
