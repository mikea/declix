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
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := pkl.LoadFromPath(context.Background(), args[0])
		if err != nil {
			panic(err)
		}

		pterm.Println()
		pterm.Println("Target: ", cfg.Target.Address)
		pterm.Println()
		pterm.Println(pterm.Gray("Checking resources..."))
		progress, err := pterm.DefaultProgressbar.WithTotal(len(cfg.Resources)).Start()
		if err != nil {
			panic(err)
		}

		progress.UpdateTitle("Dialing in...")
		executor, err := impl.SshExecutor(*cfg.Target)
		if err != nil {
			panic(err)
		}
		defer executor.Close()

		resources := make([]interfaces.Resource, len(cfg.Resources))
		for i, res := range cfg.Resources {
			resources[i] = impl.CreateResource(res)
		}

		states := make([]interfaces.ResouceStatus, len(resources))
		for i, res := range resources {
			progress.UpdateTitle(res.Id())
			progress.Increment()
			states[i] = res.DetermineStatus(executor)
		}
		progress.Stop()
		pterm.Println()

		tableData := make(pterm.TableData, len(states)+1)
		tableData[0] = []string{"Resource Id", "Status", "Expected"}
		for i, state := range states {
			tableData[i+1] = []string{resources[i].Id(), state.StyledString(), resources[i].ExpectedStatusStyledString()}
		}

		pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
