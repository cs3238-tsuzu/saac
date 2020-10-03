package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const healthEP = "127.0.0.1:16274"

func child() {
	commands := strings.Split(os.Getenv("CHILD_COMMAND"), ",")
	cmd := exec.Command(commands[0], commands[1:]...)

	if err := cmd.Start(); err != nil {
		log.Fatalf("failed to start child: %+v", err)
	}

	mux := http.NewServeMux()

	http.ListenAndServe(healthEP, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

	}))
}
