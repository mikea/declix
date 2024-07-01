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

	statuses := make([]PackageStatus, len(cfg.Packages))
	for i, pkg := range cfg.Packages {
		progress.UpdateTitle(fmt.Sprint("package:", pkg.Name))
		progress.Increment()
		statuses[i] = processPackage(client, pkg)
		statuses[i].Package = pkg
	}
	progress.Stop()

	for _, status := range statuses {
		printStatus(status)
	}

	installPackages(client, statuses)
}

func installPackages(client *ssh.Client, statuses []PackageStatus) {
	var toInstall []string
	for _, status := range statuses {
		if status.Status == ToCreate {
			toInstall = append(toInstall, status.Package.Name)
		}
	}

	output, err := execute(client, fmt.Sprintf("sudo -S apt-get install -y --no-upgrade --no-install-recommends %s", strings.Join(toInstall, " ")))
	if err != nil {
		pterm.Println(string(output))
		log.Fatalf("Failed to execute apt install: %s", err)
	}
}

func execute(client *ssh.Client, cmd string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

func printStatus(status PackageStatus) {
	switch status.Status {

	case Error:
		pterm.Println(pterm.BgRed.Sprint("E package:", status.Package.Name))
	case Ok:
		pterm.Printfln("  package:%s %s", status.Package.Name, status.Version)
	case ToCreate:
		pterm.Println(pterm.Green("+ package:", status.Package.Name))
	case ToDelete:
		pterm.Println(pterm.Red("- package:", status.Package.Name))
	case ToFix:
		pterm.Println(pterm.Yellow("~ package:", status.Package.Name))
	default:
		panic(fmt.Sprintf("unexpected main.Status: %#v", status.Status))
	}
}

type Status int

const (
	Ok Status = iota
	ToCreate
	ToDelete
	ToFix
	Error
)

type PackageStatus struct {
	Package *system.Package
	Status  Status
	Version string
}

func processPackage(client *ssh.Client, pkg *system.Package) PackageStatus {
	output, err := execute(client, fmt.Sprintf("dpkg-query -W -f='${Version}' %s", pkg.Name))
	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return PackageStatus{
				Status: Error,
			}
		}

		if e.ExitStatus() == 1 {
			return PackageStatus{
				Status: ToCreate,
			}
		} else {
			log.Fatalf("Exit error: %v", e)
			return PackageStatus{
				Status: Error,
			}
		}
	}

	return PackageStatus{
		Status:  Ok,
		Version: string(output),
	}
}
