package main

import (
	"net/http"
	"net/http/cgi"
	"os"
)

func server() {
	cgi.Serve(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

	}))
	os.Stdout.Close()
}
