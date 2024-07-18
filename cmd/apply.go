/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"mikea/declix/impl"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"mikea/declix/target"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:     "apply",
	GroupID: "main",
	Short:   "Apply all actions",
	Long:    `Apply all necessary actions to bring system to desired state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		targetPkl, err := target.LoadFromPath(context.Background(), targetFile)
		if err != nil {
			return err
		}
		target := targetPkl.Target
		pterm.Println("Target: ", target.Address)

		resourcesPkl, err := resources.LoadFromPath(context.Background(), resourcesFile)
		if err != nil {
			return err
		}

		resources := impl.CreateResources(resourcesPkl.Resources)

		pterm.FgGray.Println("Checking...")
		progress, err := pterm.DefaultProgressbar.WithTotal(len(resources)).WithRemoveWhenDone(true).Start()
		if err != nil {
			return err
		}

		progress.UpdateTitle("Connecting...")
		executor, err := impl.SshExecutor(*target)
		if err != nil {
			return err
		}
		defer executor.Close()

		states, expectedStates, errors := impl.DetermineStates(resources, executor, *progress)
		progress.Stop()

		actions := make([]interfaces.Action, len(resources))
		totalActions := 0
		for i, res := range resources {
			if errors[i] != nil {
				continue
			}
			actions[i], errors[i] = res.DetermineAction(executor, states[i], expectedStates[i])
			if actions[i] != nil {
				totalActions += 1
			}
		}

		pterm.FgGray.Println("Applying...")
		progress, err = pterm.DefaultProgressbar.WithTotal(totalActions).WithRemoveWhenDone(true).Start()
		if err != nil {
			return err
		}

		for i, action := range actions {
			if actions[i] == nil {
				continue
			}
			progress.UpdateTitle(action.StyledString(resources[i]))
			progress.Increment()

			err = resources[i].RunAction(executor, actions[i], states[i], expectedStates[i])
			if err != nil {
				pterm.Println(pterm.BgRed.Sprint("E ", actions[i].StyledString(resources[i])), err)
			} else {
				pterm.Println(pterm.FgGreen.Sprint("\u2713"), actions[i].StyledString(resources[i]))
			}
		}
		progress.Stop()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
