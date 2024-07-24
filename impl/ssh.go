package impl

import (
	"context"
	"fmt"
	"io"
	"log"
	"mikea/declix/interfaces"
	"mikea/declix/target"
	"os"
	"strings"

	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

type sshExecutor struct {
	pkl    target.SshConfig
	client *ssh.Client
}

// MkTemp implements interfaces.CommandExecutor.
func (executor sshExecutor) MkTemp() (string, error) {
	tmp, err := executor.Run("mktemp")
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(tmp, "\n"), nil
}

// UploadTemp implements interfaces.CommandExecutor.
func (executor sshExecutor) UploadTemp(content io.Reader, size int64) (string, error) {
	tmp, err := executor.MkTemp()
	if err != nil {
		return "", err
	}
	err = executor.Upload(content, tmp, "0644", size)
	if err != nil {
		return "", fmt.Errorf("error uploading file: %w", err)
	}
	return tmp, nil
}

// Upload implements interfaces.CommandExecutor.
func (s sshExecutor) Upload(content io.Reader, remotePath string, permissions string, size int64) error {
	client, err := scp.NewClientBySSH(s.client)
	if err != nil {
		return fmt.Errorf("can't create scp client: %w", err)
	}
	defer client.Close()

	if size >= 0 {
		return client.Copy(context.Background(), content, remotePath, permissions, size)
	} else {
		tmp, err := os.CreateTemp("", "declix")
		if err != nil {
			return err
		}
		tmpName := tmp.Name()
		defer os.Remove(tmpName)
		defer tmp.Close()
		if size, err = io.Copy(tmp, content); err != nil {
			return err
		}

		if err = tmp.Close(); err != nil {
			return err
		}
		if tmp, err = os.Open(tmpName); err != nil {
			return err
		}

		err = client.Copy(context.Background(), tmp, remotePath, permissions, size)
		if err != nil {
			return err
		}
		return err
	}
}

func SshExecutor(pkl target.SshConfig) (interfaces.CommandExecutor, error) {
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

// Run implements interfaces.CommandExecutor.
func (s sshExecutor) Run(command string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("error creating new session: %w", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return string(output), fmt.Errorf("%w %s", err, output)
	}
	return string(output), err
}

func (s sshExecutor) Close() error {
	return s.client.Close()
}
