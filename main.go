package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mikea/declix/system"
	"os"

	"golang.org/x/crypto/ssh"
)

var file = flag.String("file", "", "File to load")
var host = flag.String("hist", "", "Host to apply the file to")

func main() {
	flag.Parse()
	cfg, err := system.LoadFromPath(context.Background(), *file)
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

	for _, pkg := range cfg.Packages {
		session, err := client.NewSession()
		if err != nil {
			log.Fatalf("Failed to create session: %s", err)
		}
		defer session.Close()

		// output, err := session.CombinedOutput(fmt.Sprintf("apt-cache policy %s", pkg.Name))
		output, err := session.CombinedOutput(fmt.Sprintf("dpkg-query -W -f=${Version} %s", pkg.Name))
		if err != nil {
			log.Fatalf("Failed to run command: %s", err)
		}

		os.Stdout.Write(output)
	}

}
