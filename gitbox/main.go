package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	rootDir := os.Getenv("GIT_PROJECT_ROOT")
	if rootDir == "" {
		fmt.Fprintln(os.Stderr, "GIT_PROJECT_ROOT not set")
		os.Exit(1)
	}
	rootDir, err := filepath.Abs(rootDir)
	check(err)

	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "AUTH_TOKEN not set")
		os.Exit(1)
	}

	if skip, _ := strconv.ParseBool(os.Getenv("SKIP_ROOT_DIR_INIT")); !skip {
		_, err = os.Stat(filepath.Join(rootDir, "HEAD"))
		if errors.Is(err, os.ErrNotExist) {
			err = exec.Command("git", "init", "--bare", rootDir).Run()
		}
		check(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	git, err := exec.LookPath("git")
	if err != nil {
		fmt.Fprintln(os.Stderr, "git not found", err)
		os.Exit(1)
	}

	cgiErr := loggerWriter{log.New(os.Stderr, "git-http-backend: ", 0)}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, reqToken, _ := r.BasicAuth()
		if reqToken != token {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		(&cgi.Handler{
			Path:   git,
			Args:   []string{"http-backend"},
			Dir:    rootDir,
			Stderr: cgiErr,
			Env: []string{
				"GIT_PROJECT_ROOT=" + rootDir,
				"GIT_HTTP_EXPORT_ALL=yes",
				"REMOTE_USER=" + user,
			},
			Logger: cgiErr.Logger,
		}).ServeHTTP(w, r)
	})

	fmt.Println("Listening on port", port)

	cert := os.Getenv("TLS_CERT")
	key := os.Getenv("TLS_KEY")
	if cert == "" || key == "" {
		err = http.ListenAndServe(":"+port, handler)
	} else {
		err = http.ListenAndServeTLS(":"+port, cert, key, handler)
	}
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type loggerWriter struct{ *log.Logger }

func (s loggerWriter) Write(p []byte) (n int, err error) {
	s.Logger.Print(string(p))
	return len(p), nil
}
