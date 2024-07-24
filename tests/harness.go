package tests

import (
	"context"
	"fmt"
	"mikea/declix/impl"
	"os"
	"time"

	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/specgen"
	. "github.com/onsi/ginkgo/v2"
)

type Harness struct {
	App         *impl.App
	sshHostPort string
	conn        context.Context
	containerID string
}

func NewHarness() *Harness {
	return &Harness{
		App: &impl.App{},
	}
}

func (h *Harness) StartTarget() (err error) {
	ctx := context.Background()

	// connect
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sock_dir + "/podman/podman.sock"
	GinkgoWriter.Println("connecting to", socket)
	h.conn, err = bindings.NewConnection(ctx, socket)
	if err != nil {
		return err
	}

	// create new container
	GinkgoWriter.Println("creating container")
	s := specgen.NewSpecGenerator("declix-test", false)
	s.PortMappings = append(s.PortMappings, types.PortMapping{
		ContainerPort: 22,
	})
	r, err := containers.CreateWithSpec(h.conn, s, &containers.CreateOptions{})
	if err != nil {
		return err
	}
	h.containerID = r.ID

	GinkgoWriter.Println("starting container", h.containerID)
	err = containers.Start(h.conn, h.containerID, &containers.StartOptions{})
	if err != nil {
		return err
	}

	inspect, err := containers.Inspect(h.conn, h.containerID, &containers.InspectOptions{})
	if err != nil {
		return err
	}
	h.sshHostPort = inspect.HostConfig.PortBindings["22/tcp"][0].HostPort

	// wait for ssh to be up
	// this doesn't work: err = nil, but ssh still doesn't work.
	// for {
	// 	conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", sshHostPort), 10*time.Millisecond)
	// 	if err == nil {
	// 		conn.Close()
	// 		break
	// 	}
	// }
	time.Sleep(1000 * time.Millisecond)

	// load target into the app
	err = h.App.LoadTargetFromText(fmt.Sprintf(`
			amends "package://declix.org/pkl@0.0.8#/target/Target.pkl"
			
			target = new SshConfig {
				user = "declix"
				address = "localhost:%s"
				privateKey = "id_rsa"
			}	
		`, h.sshHostPort))
	if err != nil {
		return err
	}

	return nil
}

func (h *Harness) CleanupTarget() (err error) {
	h.App.Dispose()
	GinkgoWriter.Println("stopping container", h.containerID)
	var timeout uint = 1
	err = containers.Stop(h.conn, h.containerID, &containers.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		return err
	}

	GinkgoWriter.Println("removing container", h.containerID)
	_, err = containers.Remove(h.conn, h.containerID, &containers.RemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}
