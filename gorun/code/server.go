package main

import "net/http"

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, World!\n"))
	}))
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
