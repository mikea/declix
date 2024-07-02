/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	apt_package "mikea/declix/impl"
	"mikea/declix/pkl"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
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
		pterm.Println("Target: ", cfg.Ssh.Address)
		pterm.Println()
		pterm.Println(pterm.Gray("Checking resources..."))
		progress, err := pterm.DefaultProgressbar.WithTotal(len(cfg.Packages)).Start()
		if err != nil {
			panic(err)
		}

		key, err := os.ReadFile(cfg.Ssh.PrivateKey)
		if err != nil {
			log.Fatalf("Unable to read private key: %v", err)
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("Unable to parse private key: %v", err)
		}

		progress.UpdateTitle("Dialing in...")
		sshConfig := &ssh.ClientConfig{
			User: cfg.Ssh.User,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		client, err := ssh.Dial("tcp", cfg.Ssh.Address, sshConfig)
		if err != nil {
			panic(err)
		}
		defer client.Close()

		states := make([]apt_package.PackageState, len(cfg.Packages))
		for i, pkg := range cfg.Packages {
			progress.UpdateTitle(fmt.Sprint("package:", pkg.Name))
			progress.Increment()
			states[i] = apt_package.Process(client, pkg)
			states[i].Package = pkg
		}
		progress.Stop()
		pterm.Println()

		tableData := make(pterm.TableData, len(states)+1)
		tableData[0] = []string{"Resource Id", "Status"}
		for i, state := range states {
			tableData[i+1] = []string{state.Package.Name, state.State.StyledString()}
		}

		pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
