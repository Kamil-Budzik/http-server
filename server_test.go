package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func TestHandleConnection(t *testing.T) {
	go func() {
		main()
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	tests := []struct {
		request      string
		expectedCode string
	}{
		{
			request:      "GET /abcdefg HTTP/1.1\r\n\r\n",
			expectedCode: "HTTP/1.1 404 Not Found",
		},
		{
			request:      "GET / HTTP/1.1\r\n\r\n",
			expectedCode: "HTTP/1.1 200 OK",
		},
	}

	for _, tt := range tests {
		conn, err := net.Dial("tcp", "localhost:4221")
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		_, err = conn.Write([]byte(tt.request))
		if err != nil {
			t.Fatalf("Failed to write to connection: %v", err)
		}

		reader := bufio.NewReader(conn)
		responseLine, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		if !strings.HasPrefix(responseLine, tt.expectedCode) {
			t.Errorf("Expected response to start with %q, but got %q", tt.expectedCode, responseLine)
		}
	}
}
