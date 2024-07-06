/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"mikea/declix/impl"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	GroupID: "main",
	Short:   "Determine status of every resource",
	Long: `Determine status of every resource. 
	
Outputs the table of the current and desired resource status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := pkl.LoadFromPath(context.Background(), args[0])
		if err != nil {
			return err
		}

		pterm.Println("Target:", cfg.Target.Address)

		resources := impl.CreateResources(cfg.Resources)

		pterm.FgGray.Println("Checking...")
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

		statuses, errors := impl.DetermineStatuses(resources, executor, *progress)
		progress.Stop()

		tableData := make(pterm.TableData, len(statuses)+1)
		tableData[0] = []string{"Resource Id", "Status", "Expected"}
		for i, status := range statuses {
			var statusString string
			if errors[i] != nil {
				statusString = pterm.BgRed.Sprint("ERROR")
			} else {
				statusString = status.StyledString(resources[i])
			}
			expected, err := resources[i].ExpectedStatusStyledString()
			if err != nil {
				return err
			}
			tableData[i+1] = []string{resources[i].Id(), statusString, expected}
		}

		pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()

		hasErrors := false
		for _, err := range errors {
			if err != nil {
				hasErrors = true
				break
			}
		}

		if hasErrors {
			pterm.Println(pterm.FgRed.Sprint("Errors:"))
			for i, err := range errors {
				if err != nil {
					pterm.Println(pterm.BgRed.Sprint(resources[i].Id()), pterm.FgRed.Sprint(err))
				}
			}

			return fmt.Errorf("errors occured")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
