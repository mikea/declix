package filesystem

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mikea/declix/content"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

// RunAction implements interfaces.Resource.
func (f FileImpl) RunAction(executor interfaces.CommandExcutor, a interfaces.Action, s interfaces.State, es interfaces.State) error {
	expected := es.(state)

	action := a.(action)
	switch action {
	case ToUpload:
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

func (f FileImpl) update(executor interfaces.CommandExcutor, status state, expectedStatus state) error {
	if status.Group != expectedStatus.Group {
		if err := f.chgrp(executor, expectedStatus); err != nil {
			return err
		}
	}
	if status.Owner != expectedStatus.Owner {
		if err := f.chown(executor, expectedStatus); err != nil {
			return err
		}
	}
	if status.Permissions != expectedStatus.Permissions {
		if err := f.chmod(executor, expectedStatus); err != nil {
			return err
		}
	}

	return nil
}

func (f FileImpl) chmod(executor interfaces.CommandExcutor, expectedStatus state) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chmod %s %s", expectedStatus.Permissions, f.Path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (f FileImpl) chown(executor interfaces.CommandExcutor, expectedStatus state) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chown %s %s", expectedStatus.Owner, f.Path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (f FileImpl) chgrp(executor interfaces.CommandExcutor, expectedStatus state) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chgrp %s %s", expectedStatus.Group, f.Path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (f FileImpl) openContent() (io.ReadCloser, int64, error) {
	return content.OpenContent(f.State.(*Present).Content)
}

func (f FileImpl) upload(executor interfaces.CommandExcutor, expectedStatus state) error {
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

	if err := f.chown(executor, expectedStatus); err != nil {
		return err
	}
	if err := f.chgrp(executor, expectedStatus); err != nil {
		return err
	}
	if err := f.chmod(executor, expectedStatus); err != nil {
		return err
	}

	return nil
}

// DetermineAction implements interfaces.Resource.
func (f FileImpl) DetermineAction(executor interfaces.CommandExcutor, s interfaces.State, es interfaces.State) (interfaces.Action, error) {
	expected := es.(state)
	state := s.(state)

	if expected.Exists {
		if state.Exists {
			if state.Sha256 != expected.Sha256 {
				return ToUpload, nil
			}
			if state.Owner != expected.Owner ||
				state.Group != expected.Group ||
				state.Permissions != expected.Permissions {
				return ToUpdate, nil
			}

			return nil, nil
		}

		return ToUpload, nil
	}

	return ToDelete, nil
}

type state struct {
	Exists      bool
	Size        int64
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
	case ToUpload:
		return pterm.FgGreen.Sprint("+", resource.Id())
	case ToUpdate:
		return pterm.FgYellow.Sprint("~", resource.Id())
	case ToDelete:
		return pterm.FgRed.Sprint("-", resource.Id())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

const (
	ToUpload action = iota
	ToUpdate
	ToDelete
)

// DetermineState implements interfaces.Resource.
func (f FileImpl) DetermineState(executor interfaces.CommandExcutor) (interfaces.State, error) {
	out, err := executor.Run(fmt.Sprintf(
		`if [ ! -f "%s" ]; then 
			echo "exists: false"; 
		else 
			echo "exists: true" &&
			read -r hash _ < <(sudo sha256sum %s) &&
			echo "sha256: $hash" &&
			stat --printf="size: %%s\nowner: %%U\ngroup: %%G\npermissions: %%a\n" %s
		fi`,
		f.Path,
		f.Path,
		f.Path,
	))
	if err != nil {
		return nil, err
	}
	state := state{}
	if err := yaml.Unmarshal([]byte(out), &state); err != nil {
		return nil, err
	}
	return state, nil
}

// Id implements interfaces.Resource.
func (f FileImpl) Id() string {
	return fmt.Sprintf("%s:%s", f.Type, f.Path)
}

// Pkl implements interfaces.Resource.
func (f FileImpl) Pkl() resources.Resource {
	return f
}

func (f FileImpl) ExpectedState() (interfaces.State, error) {
	switch s := f.State.(type) {
	case *Missing:
		return state{
			Exists: false,
		}, nil
	case *Present:
		content, size, err := f.openContent()
		if err != nil {
			return nil, err
		}
		defer content.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, content); err != nil {
			return nil, err
		}
		sha256 := fmt.Sprintf("%x", string(hasher.Sum(nil)))

		return state{
			Exists:      true,
			Size:        size,
			Sha256:      sha256,
			Owner:       s.Owner,
			Group:       s.Group,
			Permissions: s.Permissions,
		}, nil

	default:
		panic(fmt.Sprintf("unsupported state %T", s))

	}

}
