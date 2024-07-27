/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mikea/declix/impl"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:     "apply",
	GroupID: "main",
	Short:   "Apply all actions",
	Long:    `Apply all necessary actions to bring system to desired state.`,
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
		if err := app.ApplyActions(); err != nil {
			return err
		}

		if app.HasErrors() {
			app.PrintErrors()
			return fmt.Errorf("there were errors applying actions")
		}

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
