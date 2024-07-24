package filesystem

import (
	"fmt"
	"mikea/declix/interfaces"
)

func chown(executor interfaces.CommandExecutor, path string, owner string) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chown %s \"%s\"", owner, path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func chmod(executor interfaces.CommandExecutor, path string, permissions string) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chmod %s \"%s\"", permissions, path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}

func chgrp(executor interfaces.CommandExecutor, path string, group string) error {
	_, err := executor.Run(fmt.Sprintf("sudo -S chgrp %s \"%s\"", group, path))
	if err != nil {
		return fmt.Errorf("error changing permissions: %w", err)
	}
	return nil
}
