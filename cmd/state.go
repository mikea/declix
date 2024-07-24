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

	tableData := make(pterm.TableData, len(app.States)+1)
	tableData[0] = []string{"Resource Id", "Current State", "Expected State"}
	for i, state := range app.States {
		var stateStr string
		var expectedStr string

		if state != nil {
			stateStr = state.StyledString(app.Resources[i])
		} else {
			stateStr = pterm.BgRed.Sprint("ERROR")
		}
		if app.Expected[i] != nil {
			expectedStr = app.Expected[i].StyledString(app.Resources[i])
		} else {
			expectedStr = pterm.BgRed.Sprint("ERROR")
		}

		tableData[i+1] = []string{app.Resources[i].GetId(), stateStr, expectedStr}
	}

	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()

	hasErrors := false
	for _, err := range app.Errors {
		if err != nil {
			hasErrors = true
			break
		}
	}

	if hasErrors {
		pterm.Println(pterm.FgRed.Sprint("Errors:"))
		for i, err := range app.Errors {
			if err != nil {
				pterm.Println(pterm.BgRed.Sprint(app.Resources[i].GetId()), pterm.FgRed.Sprint(err))
			}
		}

		return fmt.Errorf("errors occured")
	}

	return nil
}
