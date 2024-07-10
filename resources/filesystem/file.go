package filesystem

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

func New(pkl File) interfaces.Resource {
	return resource{pkl: pkl}
}

type resource struct {
	pkl File
}

// RunAction implements interfaces.Resource.
func (r resource) RunAction(executor interfaces.CommandExcutor, a interfaces.Action, s interfaces.Status) error {
	expectedStatus, err := r.ExpectedStatus()
	if err != nil {
		return err
	}

	action := a.(action)
	switch action {
	case ToUpload:
		return r.upload(executor, *expectedStatus)
	case ToDelete:
		panic("not implemented")
	case ToUpdate:
		return r.update(executor, s.(status), *expectedStatus)
	default:
		panic(fmt.Sprintf("unexpected filesystem.action: %#v", action))
	}
}

func (r resource) update(executor interfaces.CommandExcutor, status status, expectedStatus status) error {
	if status.Group != expectedStatus.Group {
		if err := r.chgrp(executor, expectedStatus); err != nil {
			return err
		}
	}
	if status.Owner != expectedStatus.Owner {
		if err := r.chown(executor, expectedStatus); err != nil {
			return err
		}
	}
	if status.Permissions != expectedStatus.Permissions {
		if err := r.chmod(executor, expectedStatus); err != nil {
			return err
		}
	}

	return nil
}

func (r resource) chmod(executor interfaces.CommandExcutor, expectedStatus status) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chmod %s %s", expectedStatus.Permissions, r.pkl.GetPath()))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (r resource) chown(executor interfaces.CommandExcutor, expectedStatus status) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chown %s %s", expectedStatus.Owner, r.pkl.GetPath()))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (r resource) chgrp(executor interfaces.CommandExcutor, expectedStatus status) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chgrp %s %s", expectedStatus.Group, r.pkl.GetPath()))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func (r resource) upload(executor interfaces.CommandExcutor, expectedStatus status) error {
	tmp, err := executor.Run("mktemp")
	if err != nil {
		return err
	}
	tmp = strings.TrimSuffix(tmp, "\n")

	content, size, err := r.GetContent()
	if err != nil {
		return err
	}
	defer content.Close()

	err = executor.Upload(content, tmp, "0644", size)
	if err != nil {
		return fmt.Errorf("error uploading file: %w", err)
	}

	_, err = executor.Run(fmt.Sprintf("sudo -S mv %s %s", tmp, r.pkl.GetPath()))
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	if err := r.chown(executor, expectedStatus); err != nil {
		return err
	}
	if err := r.chgrp(executor, expectedStatus); err != nil {
		return err
	}
	if err := r.chmod(executor, expectedStatus); err != nil {
		return err
	}

	return nil
}

// DetermineAction implements interfaces.Resource.
func (r resource) DetermineAction(executor interfaces.CommandExcutor, s interfaces.Status) (interfaces.Action, error) {
	status := s.(status)
	expectedStatus, err := r.ExpectedStatus()
	if err != nil {
		return nil, err
	}

	if expectedStatus.Exists {
		if status.Exists {
			if status.Sha256 != expectedStatus.Sha256 {
				return ToUpload, nil
			}
			if status.Owner != expectedStatus.Owner ||
				status.Group != expectedStatus.Group ||
				status.Permissions != expectedStatus.Permissions {
				return ToUpdate, nil
			}

			return nil, nil
		}

		return ToUpload, nil
	} else {
		panic("not implemented")
	}
}

type noCloseReadCloser struct {
	reader io.Reader
}

// Close implements io.ReadCloser.
func (noCloseReadCloser) Close() error {
	return nil
}

// Read implements io.ReadCloser.
func (r noCloseReadCloser) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r resource) GetContent() (io.ReadCloser, int64, error) {
	switch c := r.pkl.GetContent().(type) {
	case *FileContent:
		f, err := os.Open(c.File)
		if err != nil {
			return nil, 0, err
		}
		stat, err := f.Stat()
		if err != nil {
			return nil, 0, err
		}
		return f, stat.Size(), nil
	case string:
		return noCloseReadCloser{strings.NewReader(c)}, int64(len(c)), nil
	default:
		panic(fmt.Sprintf("unsupported content %T", c))
	}
}

type status struct {
	Exists      bool
	Size        int64
	Sha256      string
	Owner       string
	Group       string
	Permissions string
}

// StyledString implements interfaces.ResouceStatus.
func (s status) StyledString(resource interfaces.Resource) string {
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

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() (string, error) {
	expectedStatus, err := r.ExpectedStatus()
	if err != nil {
		return "", err
	}

	return pterm.FgGreen.Sprint(
		expectedStatus.Sha256[:8],
		" ",
		expectedStatus.Owner,
		":",
		expectedStatus.Group,
		" ",
		expectedStatus.Permissions,
	), nil
}

// DetermineStatus implements interfaces.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) (interfaces.Status, error) {
	out, err := executor.Run(fmt.Sprintf(
		`if [ ! -f "%s" ]; then 
			echo "exists: false"; 
		else 
			echo "exists: true" &&
			read -r hash _ < <(sudo sha256sum %s) &&
			echo "sha256: $hash" &&
			stat --printf="size: %%s\nowner: %%U\ngroup: %%G\npermissions: %%a\n" %s
		fi`,
		r.pkl.GetPath(),
		r.pkl.GetPath(),
		r.pkl.GetPath(),
	))
	if err != nil {
		return nil, err
	}
	status := status{}
	yaml.Unmarshal([]byte(out), &status)
	return status, nil
}

// Id implements interfaces.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetPath())
}

// Pkl implements interfaces.Resource.
func (r resource) Pkl() resources.Resource {
	return r.pkl
}

func (r resource) ExpectedStatus() (*status, error) {
	content, size, err := r.GetContent()
	if err != nil {
		return nil, err
	}
	defer content.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, content); err != nil {
		return nil, err
	}
	sha256 := fmt.Sprintf("%x", string(hasher.Sum(nil)))

	return &status{
		Exists:      true,
		Size:        size,
		Sha256:      sha256,
		Owner:       r.pkl.GetOwner(),
		Group:       r.pkl.GetGroup(),
		Permissions: r.pkl.GetPermissions(),
	}, nil
}
