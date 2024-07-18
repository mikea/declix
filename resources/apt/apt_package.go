package apt

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"mikea/declix/resources/dpkg"

	"github.com/pterm/pterm"
)

type resource struct {
	pkl Package
}

// RunAction implements interfaces.Resource.
func (r resource) RunAction(executor interfaces.CommandExcutor, a interfaces.Action, status interfaces.Status) error {
	switch a.(action) {
	case ToInstall:
		cmd := fmt.Sprintf("sudo -S apt-get install -y --no-upgrade --no-install-recommends %s", r.pkl.GetName())
		if r.pkl.GetUpdateBeforeInstall() {
			cmd = fmt.Sprintf("sudo -S apt-get update && %s", cmd)
		}
		_, err := executor.Run(cmd)
		return err
	case ToRemove:
		_, err := executor.Run(fmt.Sprintf("sudo -S apt-get remove -y %s", r.pkl.GetName()))
		return err
	}
	panic("unexpected apt_package.action")
}

// DetermineAction implements interfaces.Resource.
func (r resource) DetermineAction(executor interfaces.CommandExcutor, s interfaces.Status) (interfaces.Action, error) {
	status := s.(dpkg.Status)

	switch r.pkl.GetStatus() {
	case "installed":
		if status.PackageState() == dpkg.Installed {
			return nil, nil
		}
		return ToInstall, nil
	case "missing":
		if status.PackageState() == dpkg.Missing {
			return nil, nil
		}
		return ToRemove, nil
	}

	panic(fmt.Sprintf("unexpected status: %#v", r.pkl.GetStatus()))
}

type action int

// StyledString implements interfaces.Action.
func (a action) StyledString(resource interfaces.Resource) string {
	switch a {
	case ToInstall:
		return pterm.FgGreen.Sprint("+", resource.Id())
	case ToRemove:
		return pterm.FgRed.Sprint("-", resource.Id())
	}
	panic(fmt.Sprintf("unexpected apt_package.action: %#v", a))
}

const (
	ToInstall action = iota
	ToRemove  action = iota
)

func New(pkl Package) interfaces.Resource {
	return resource{pkl: pkl}
}

// ExpectedStatusStyledString implements interfaces.Resource.
func (r resource) ExpectedStatusStyledString() (string, error) {
	switch r.pkl.GetStatus() {
	case "installed":
		return pterm.FgGreen.Sprint("installed"), nil
	case "missing":
		return pterm.FgRed.Sprint("missing"), nil
	}

	panic(fmt.Sprintf("unexpected status: %#v", r.pkl.GetStatus()))
}

// DetermineStatus implements impl.Resource.
func (r resource) DetermineStatus(executor interfaces.CommandExcutor) (interfaces.Status, error) {
	name := r.pkl.GetName()
	return dpkg.DeterminePackageStatus(executor, name)
}

// Id implements impl.Resource.
func (r resource) Id() string {
	return fmt.Sprintf("%s:%s", r.pkl.GetType(), r.pkl.GetName())
}

// Pkl implements impl.Resource.
func (r resource) Pkl() resources.Resource {
	return r.pkl
}
