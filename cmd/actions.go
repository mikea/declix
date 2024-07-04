/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"mikea/declix/impl"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:     "actions",
	GroupID: "main",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short:   "List all actions to perform",
	Long: `Determine and list all actions to perform to bring the system
th the desired state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := pkl.LoadFromPath(context.Background(), args[0])
		if err != nil {
			return err
		}

		pterm.Println("Target: ", cfg.Target.Address)

		resources := impl.CreateResources(cfg.Resources)

		progress, err := pterm.DefaultProgressbar.WithTotal(len(resources)).WithRemoveWhenDone(true).Start()
		if err != nil {
			return err
		}

		progress.UpdateTitle("Connecting...")
		executor, err := impl.SshExecutor(*cfg.Target)
		if err != nil {
			return err
		}
		defer executor.Close()

		pterm.FgGray.Println("Checking...")
		statuses, errors := impl.DetermineStatuses(resources, executor, *progress)
		progress.Stop()

		actions := make([]interfaces.Action, len(resources))
		for i, res := range resources {
			if errors[i] == nil {
				actions[i], errors[i] = res.DetermineAction(executor, statuses[i])
			}
		}

		for i, action := range actions {
			if errors[i] != nil {
				pterm.BgRed.Sprint(resources[i].Id())
			} else if action != nil {
				pterm.Println(action.StyledString(resources[i]))
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)

}
