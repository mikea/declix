package systemd

import (
	"fmt"
	"mikea/declix/interfaces"
)

func (f *UnitFileImpl) RunAction(executor interfaces.CommandExecutor, a interfaces.Action, s interfaces.State, e interfaces.State) error {
	if err := f.FileImpl.RunAction(executor, a, s, e); err != nil {
		return err
	}

	if f.AfterAction.DaemonReload {
		if _, err := executor.Run("sudo systemctl daemon-reload"); err != nil {
			return err
		}
	}

	if f.AfterAction.ReloadOrRestart != nil {
		if _, err := executor.Run(fmt.Sprintf("sudo systemctl reload-or-restart %s", *f.AfterAction.ReloadOrRestart)); err != nil {
			return err
		}
	}
	if f.AfterAction.Reload != nil {
		if _, err := executor.Run(fmt.Sprintf("sudo systemctl reload %s", *f.AfterAction.Reload)); err != nil {
			return err
		}
	}
	if f.AfterAction.Restart != nil {
		if _, err := executor.Run(fmt.Sprintf("sudo systemctl restart %s", *f.AfterAction.Restart)); err != nil {
			return err
		}
	}
	return nil
}
