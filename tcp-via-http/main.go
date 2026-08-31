package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s <listen> <http-proxy-addr> <target>\n", os.Args[0])
		os.Exit(2)
	}
	listen, proxy, target := os.Args[1], os.Args[2], os.Args[3]
	if !strings.Contains(listen, ":") {
		listen = ":" + listen
	}
	if _, _, err := net.SplitHostPort(target); err != nil {
		log.Fatalf("bad target %q: %v", target, err)
	}

	ln, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s -> %s via http proxy %s", ln.Addr(), target, proxy)

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handle(c, proxy, target)
	}
}

func handle(client net.Conn, proxy, target string) {
	defer client.Close()

	up, err := net.Dial("tcp", proxy)
	if err != nil {
		log.Printf("dial proxy: %v", err)
		return
	}
	defer up.Close()

	br, err := httpConnect(up, target)
	if err != nil {
		log.Printf("CONNECT %s: %v", target, err)
		return
	}

	done := make(chan struct{}, 2)
	go func() { // proxy -> client, starting with anything already buffered
		io.Copy(client, br)
		if c, ok := client.(*net.TCPConn); ok {
			c.CloseWrite()
		}
		done <- struct{}{}
	}()
	go func() {
		io.Copy(up, client)
		if c, ok := up.(*net.TCPConn); ok {
			c.CloseWrite()
		}
		done <- struct{}{}
	}()
	<-done
	<-done
}

// httpConnect sends CONNECT host:port and returns a reader that holds any
// bytes the proxy already sent after the response headers.
func httpConnect(c net.Conn, target string) (*bufio.Reader, error) {
	req := "CONNECT " + target + " HTTP/1.1\r\nHost: " + target + "\r\n\r\n"
	if _, err := io.WriteString(c, req); err != nil {
		return nil, err
	}
	br := bufio.NewReader(c)
	resp, err := http.ReadResponse(br, &http.Request{Method: "CONNECT"})
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("proxy returned %s", resp.Status)
	}
	return br, nil
}
