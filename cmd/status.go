/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"mikea/declix/impl"
	"mikea/declix/resources"
	"mikea/declix/target"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Args:    cobra.ExactArgs(0),
	GroupID: "main",
	Short:   "Determine status of every resource",
	Long: `Determine status of every resource. 
	
Outputs the table of the current and desired resource status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func executeStatus() error {
	targetPkl, err := target.LoadFromPath(context.Background(), targetFile)
	if err != nil {
		return fmt.Errorf("loading target file: %w", err)
	}
	target := targetPkl.Target
	pterm.Printfln("Target: %v", target)

	resourcesPkl, err := resources.LoadFromPath(context.Background(), resourcesFile)
	if err != nil {
		return fmt.Errorf("loading resources file: %w", err)
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
}
