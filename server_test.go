package main

import (
	"net"
	"testing"
	"time"
)

func TestHandleConnection(t *testing.T) {
	// Start the server in a separate goroutine
	go main()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "localhost:4221")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\n"))
	if err != nil {
		t.Fatalf("Failed to write request: %v", err)
	}

	// Read the response
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if string(buf[:17]) != "HTTP/1.1 200 OK\r\n" {
		t.Errorf("Expected response to start with 'HTTP/1.1 200 OK', but got %s", string(buf[:17]))
	}
}
