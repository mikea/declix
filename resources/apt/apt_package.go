package apt

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources/dpkg"
	"mikea/declix/resources/dpkg/state"
)

// RunAction implements interfaces.Resource.
func (p PackageImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, es interfaces.State) error {
	switch a.(dpkg.Action) {
	case dpkg.ToInstall:
		cmd := fmt.Sprintf("sudo -S apt-get install -y --no-upgrade --no-install-recommends %s", p.Name)
		if p.UpdateBeforeInstall {
			cmd = fmt.Sprintf("sudo -S apt-get update && %s", cmd)
		}
		_, err := executor.Run(cmd)
		return err
	case dpkg.ToRemove:
		_, err := executor.Run(fmt.Sprintf("sudo -S apt-get remove -y %s", p.Name))
		return err
	}
	panic(fmt.Sprintf("unexpected action: %#v", a))
}

// DetermineAction implements interfaces.Resource.
func (p PackageImpl) DetermineAction(s interfaces.State, es interfaces.State) (interfaces.Action, error) {
	return dpkg.DetermineAction(s.(dpkg.State), es.(dpkg.State))
}

// ExpectedState implements interfaces.Resource.
func (p PackageImpl) ExpectedState() (interfaces.State, error) {
	return dpkg.State{
		Present: p.State == state.Installed,
		Version: "",
	}, nil
}

// DetermineState implements impl.Resource.
func (p PackageImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	name := p.Name
	return dpkg.DeterminePackageState(executor, name)
}
