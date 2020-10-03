package main

import (
	"os"
	"os/exec"
)

func Fork() error {
	cmd := exec.Command(os.Args[0])
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "CHILD=1")

	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Process.Release(); err != nil {
		return err
	}

	return nil
}

func main() {
	if os.Getenv("CHILD") != "" {
		child()

		return
	}

	server()
}
