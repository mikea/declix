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

		pterm.Println()
		pterm.Println("Target: ", cfg.Target.Address)
		pterm.Println()

		resources := impl.CreateResources(cfg.Resources)

		progress, err := pterm.DefaultProgressbar.WithTotal(len(resources)).WithRemoveWhenDone(true).Start()
		if err != nil {
			return err
		}

		progress.UpdateTitle("Dialing in...")
		executor, err := impl.SshExecutor(*cfg.Target)
		if err != nil {
			return err
		}
		defer executor.Close()

		states := make([]interfaces.Status, len(resources))
		for i, res := range resources {
			progress.UpdateTitle(res.Id())
			progress.Increment()
			states[i] = res.DetermineStatus(executor)
		}
		progress.Stop()

		tableData := make(pterm.TableData, len(states)+1)
		tableData[0] = []string{"Resource Id", "Status", "Expected"}
		for i, state := range states {
			tableData[i+1] = []string{resources[i].Id(), state.StyledString(resources[i]), resources[i].ExpectedStatusStyledString()}
		}

		pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
