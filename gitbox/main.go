package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cgi"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	rootDir := os.Getenv("GIT_PROJECT_ROOT")
	if rootDir == "" {
		fmt.Fprintln(os.Stderr, "GIT_PROJECT_ROOT not set")
		os.Exit(1)
	}

	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "AUTH_TOKEN not set")
		os.Exit(1)
	}

	_, err := os.Stat(filepath.Join(rootDir, "git-daemon-export-ok"))
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Initializing git repository")
		err = exec.Command("sh", "-ceu", `
			mkdir -p "$GIT_PROJECT_ROOT"
			git init --bare "$GIT_PROJECT_ROOT"
			cd "$GIT_PROJECT_ROOT"
			touch git-daemon-export-ok
			git config http.receivepack true
			git config http.uploadpack true
		`).Run()
	}
	check(err)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	git, err := exec.LookPath("git")
	if err != nil {
		fmt.Fprintln(os.Stderr, "git not found", err)
		os.Exit(1)
	}

	fmt.Println("Listening on port", port)

	err = http.ListenAndServe(":"+port, authMiddleware(token, &cgi.Handler{
		Path:       git,
		Args:       []string{"http-backend"},
		InheritEnv: []string{"GIT_PROJECT_ROOT"},
	}))
	check(err)
}

func authMiddleware(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, reqToken, ok := r.BasicAuth(); !ok || reqToken != token {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
