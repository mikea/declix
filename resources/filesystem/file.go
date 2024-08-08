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

func (f *FileImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, e interfaces.State) error {
	action := a.(action)
	switch action {
	case ToCreate:
		return f.upload(executor, e.(*FilePresentImpl))
	case ToDelete:
		_, err := executor.Run(fmt.Sprintf("sudo rm -f %s", f.Path))
		if err != nil {
			return err
		}
		return nil
	case ToUpdate:
		return f.update(executor, s.(*FilePresentImpl), e.(*FilePresentImpl))
	default:
		panic(fmt.Sprintf("unexpected filesystem.action: %#v", action))
	}
}

func (f *FileImpl) update(executor interfaces.CommandExecutor, current *FilePresentImpl, expected *FilePresentImpl) error {
	contentEqual, err := content.Equal(current.Content, expected.Content)
	if err != nil {
		return err
	}
	if !contentEqual {
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

func (f *FileImpl) upload(executor interfaces.CommandExecutor, expected *FilePresentImpl) error {
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
func (f *FileImpl) DetermineAction(s interfaces.State, e interfaces.State) (interfaces.Action, error) {
	switch expected := e.(type) {
	case *resources.Missing:
		if _, ok := s.(*resources.Missing); ok {
			return nil, nil
		}
		return ToDelete, nil
	case *FilePresentImpl:
		if current, ok := s.(*FilePresentImpl); ok {
			contentEqual, err := content.Equal(current.Content, expected.Content)
			if err != nil {
				return nil, err
			}
			if !contentEqual ||
				expected.Owner != current.Owner ||
				expected.Group != current.Group ||
				expected.Permissions != current.Permissions {
				return ToUpdate, nil
			}
			return nil, nil
		}
		return ToCreate, nil
	}

	panic(fmt.Sprintf("wrong state %T", e))
}

func (state *FilePresentImpl) GetStyledString() string {
	return pterm.FgGreen.Sprint(content.CachedSha256(state.Content)[:8], " ", state.Owner, ":", state.Group, " ", state.Permissions)
}

type action int

func (a action) GetStyledString(resource interfaces.Resource) string {
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

// todo: can we use pkl?
type fileStateOutput struct {
	Present     bool
	Owner       string
	Group       string
	Permissions string
	Sha256      string
}

// DetermineState implements interfaces.Resource.
func (f *FileImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	out, err := executor.Run(f.StateCmd)
	if err != nil {
		return nil, err
	}
	output := fileStateOutput{}
	if err := yaml.Unmarshal([]byte(out), &output); err != nil {
		return nil, err
	}

	if !output.Present {
		return &resources.Missing{}, nil
	}

	return &FilePresentImpl{
		Content:     &content.Hashed{Sha256: output.Sha256},
		Owner:       output.Owner,
		Group:       output.Group,
		Permissions: output.Permissions,
	}, nil
}

func (f *FileImpl) ExpectedState() (interfaces.State, error) { return f.State.(interfaces.State), nil }
