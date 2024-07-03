package impl

import (
	"log"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"
	"os"

	"golang.org/x/crypto/ssh"
)

type sshExecutor struct {
	pkl    pkl.SshConfig
	client *ssh.Client
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
