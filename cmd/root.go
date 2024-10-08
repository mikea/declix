package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var targetFile string
var resourcesFile string
var version = "DEV"

var rootCmd = &cobra.Command{
	Use:          "declix",
	Short:        "Declarative Linux",
	SilenceUsage: true,
	Version:      version,
	Long: `Declarative Linux.

Declix is a CLI tool to manage Linux systems in a declartive way.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(&cobra.Group{ID: "main", Title: "Main Commands:"})

	rootCmd.PersistentFlags().StringVarP(&targetFile, "target", "t", "", "target .pkl description file")
	rootCmd.PersistentFlags().StringVarP(&resourcesFile, "resources", "r", "", "resources .pkl description file")
}
