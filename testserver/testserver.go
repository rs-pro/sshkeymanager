package testserver

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type Server struct {
	Command *exec.Cmd
}

func IsRunning() bool {
	cmd := exec.Command("docker", "compose", "ps", "-q", "ssh")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	containerID := strings.TrimSpace(string(out))
	if containerID == "" {
		return false
	}
	// Check if container is actually running
	inspectCmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerID)
	inspectOut, err := inspectCmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(inspectOut)) == "true"
}

func Start() *Server {
	if IsRunning() {
		log.Println("test server: SSH container already running, skipping start")
		return &Server{Command: nil}
	}

	log.Println("test server: starting SSH container")
	cmd := exec.Command("docker", "compose", "up", "-d", "--force-recreate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return &Server{Command: cmd}
}

func (s *Server) Stop() {
	if s.Command == nil {
		log.Println("test server: container was already running, not stopping")
		return
	}
	stopCmd := exec.Command("docker", "compose", "down")
	stopCmd.Stdout = os.Stdout
	stopCmd.Stderr = os.Stderr
	if err := stopCmd.Run(); err != nil {
		log.Println("failed to stop containers:", err)
	}
}
