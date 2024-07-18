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
	Use:     "state",
	Args:    cobra.ExactArgs(0),
	GroupID: "main",
	Short:   "Determine state of every resource",
	Long: `Determine state of every resource. 
	
Outputs the table of the current and desired resource state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeState()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func executeState() error {
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

	states, expectedStates, errors := impl.DetermineStates(resources, executor, *progress)
	progress.Stop()

	tableData := make(pterm.TableData, len(states)+1)
	tableData[0] = []string{"Resource Id", "Current State", "Expected State"}
	for i, state := range states {
		var stateStr string
		var expectedStr string

		if state != nil {
			stateStr = state.StyledString(resources[i])
		} else {
			stateStr = pterm.BgRed.Sprint("ERROR")
		}
		if expectedStates[i] != nil {
			expectedStr = expectedStates[i].StyledString(resources[i])
		} else {
			expectedStr = pterm.BgRed.Sprint("ERROR")
		}

		tableData[i+1] = []string{resources[i].Id(), stateStr, expectedStr}
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
