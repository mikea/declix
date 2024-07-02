package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mikea/declix/system"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
)

var file = flag.String("file", "", "File to load")

func main() {
	flag.Parse()
	cfg, err := system.LoadFromPath(context.Background(), *file)
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

	states := make([]PackageState, len(cfg.Packages))
	for i, pkg := range cfg.Packages {
		progress.UpdateTitle(fmt.Sprint("package:", pkg.Name))
		progress.Increment()
		states[i] = processPackage(client, pkg)
		states[i].Package = pkg
	}
	progress.Stop()

	actions := make([]PackageAction, len(states))
	for i, state := range states {
		switch state.State {
		case ErrorState:
			actions[i] = PackageAction{
				Package: state.Package,
				Action:  Error,
			}
		case Installed:
			if cfg.Packages[i].Status == "installed" {
				actions[i] = PackageAction{
					Package: state.Package,
					Version: state.Version,
					Action:  Nothing,
				}
			} else {
				actions[i] = PackageAction{
					Package: state.Package,
					Version: state.Version,
					Action:  ToDelete,
				}
			}
		case Missing:
			if cfg.Packages[i].Status == "installed" {
				actions[i] = PackageAction{
					Package: state.Package,
					Action:  ToInstall,
				}
			} else {
				actions[i] = PackageAction{
					Package: state.Package,
					Action:  Nothing,
				}
			}
		default:
			panic(fmt.Sprintf("unexpected main.State: %#v", state.State))
		}
	}

	for _, action := range actions {
		printAction(action)
	}

	installPackages(client, actions)
}

func installPackages(client *ssh.Client, actions []PackageAction) {
	var toInstall []string
	for _, status := range actions {
		if status.Action == ToInstall {
			toInstall = append(toInstall, status.Package.Name)
		}
	}

	output, err := execute(client, fmt.Sprintf("sudo -S apt-get install -y --no-upgrade --no-install-recommends %s", strings.Join(toInstall, " ")))
	if err != nil {
		pterm.Println(string(output))
		log.Fatalf("Failed to execute apt install: %s", err)
	}
}

func printAction(action PackageAction) {
	switch action.Action {

	case Error:
		pterm.Println(pterm.BgRed.Sprint("E package:", action.Package.Name))
	case Nothing:
		pterm.Printfln("  package:%s %s", action.Package.Name, action.Version)
	case ToInstall:
		pterm.Println(pterm.Green("+ package:", action.Package.Name))
	case ToDelete:
		pterm.Println(pterm.Red("- package:", action.Package.Name))
	default:
		panic(fmt.Sprintf("unexpected main.Action: %#v", action.Action))
	}
}

type Action int

const (
	Nothing Action = iota
	ToInstall
	ToDelete
	Error
)

type PackageAction struct {
	Package *system.Package
	Action  Action
	Version string
}
