package impl

import (
	"context"
	"fmt"
	"log"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

type sshExecutor struct {
	pkl    pkl.SshConfig
	client *ssh.Client
}

// Upload implements interfaces.CommandExcutor.
func (s sshExecutor) Upload(file os.File, remotePath string, permissions string) error {
	client, err := scp.NewClientBySSH(s.client)
	if err != nil {
		return fmt.Errorf("can't create scp client: %w", err)
	}
	defer client.Close()

	return client.CopyFromFile(context.Background(), file, remotePath, permissions)
}

func SshExecutor(pkl pkl.SshConfig) (interfaces.CommandExcutor, error) {
	key, err := os.ReadFile(pkl.PrivateKey)
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: pkl.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", pkl.Address, sshConfig)
	if err != nil {
		return nil, err
	}

	return sshExecutor{pkl: pkl, client: client}, nil
}

// Run implements interfaces.CommandExcutor.
func (s sshExecutor) Run(command string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (s sshExecutor) Close() error {
	return s.client.Close()
}
