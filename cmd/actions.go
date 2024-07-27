/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mikea/declix/impl"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:     "actions",
	GroupID: "main",
	Args:    cobra.ExactArgs(0),
	Short:   "List all actions to perform",
	Long: `Determine and list all actions to perform to bring the system
th the desired state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := impl.App{}

		if err := app.LoadTarget(targetFile); err != nil {
			return err
		}
		if err := app.LoadResources(resourcesFile); err != nil {
			return err
		}
		if err := app.DetermineStates(); err != nil {
			return err
		}
		if err := app.DetermineActions(); err != nil {
			return err
		}

		for _, r := range app.Resources {
			if r.Action != nil {
				pterm.Println(r.Action.StyledString(r.Resource))
			}
		}

		if app.HasErrors() {
			app.PrintErrors()
			return fmt.Errorf("there were errors determining actions")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)

}
