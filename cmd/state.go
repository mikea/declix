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

	tableData := make(pterm.TableData, len(app.Resources)+1)
	tableData[0] = []string{"Resource Id", "Current State", "Expected State"}
	for i, r := range app.Resources {
		var currentStr string
		var expectedStr string

		if r.Current != nil {
			currentStr = r.Current.GetStyledString()
		} else {
			currentStr = pterm.BgRed.Sprint("ERROR")
		}
		if r.Expected != nil {
			expectedStr = r.Expected.GetStyledString()
		} else {
			expectedStr = pterm.BgRed.Sprint("ERROR")
		}

		tableData[i+1] = []string{r.Resource.GetId(), currentStr, expectedStr}
	}

	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()

	if app.HasErrors() {
		app.PrintErrors()
		return fmt.Errorf("errors occured")
	}

	return nil
}
