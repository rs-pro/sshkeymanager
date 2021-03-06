package testserver

import (
	"log"
	"os"
	"os/exec"
)

type Server struct {
	Command *exec.Cmd
}

func Start() *Server {
	//cmd := exec.Command("docker-compose", "up", "--force-recreate", "--build")
	cmd := exec.Command("docker-compose", "up", "--force-recreate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go func() {
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}()
	return &Server{Command: cmd}
}

func (s *Server) Stop() {
	if err := s.Command.Process.Kill(); err != nil {
		log.Fatal("failed to kill process: ", err)
	}
}
