package apt_package

import (
	"fmt"
	"log"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
	"golang.org/x/crypto/ssh"
)

type State int

const (
	ErrorState State = iota
	Missing
	Installed
)

func (s State) StyledString() string {
	switch s {

	case ErrorState:
		return pterm.BgRed.Sprint("ERROR")
	case Installed:
		return pterm.FgGreen.Sprint("installed")
	case Missing:
		return pterm.FgRed.Sprint("missing")
	}

	panic(fmt.Sprintf("unexpected apt_package.State: %#v", s))
}

type PackageState struct {
	Package *pkl.Package
	State   State
	Version string
}

func Process(client *ssh.Client, pkg *pkl.Package) PackageState {
	output, err := execute(client, fmt.Sprintf("dpkg-query -W -f='${Version}' %s", pkg.Name))
	if err != nil {
		e, ok := err.(*ssh.ExitError)
		if !ok {
			return PackageState{
				State: ErrorState,
			}
		}

		if e.ExitStatus() == 1 {
			return PackageState{
				State: Missing,
			}
		} else {
			log.Fatalf("Exit error: %v", e)
			return PackageState{
				State: ErrorState,
			}
		}
	}

	return PackageState{
		State:   Installed,
		Version: string(output),
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
