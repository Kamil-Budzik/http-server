package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestHandleConnection(t *testing.T) {
	// Start the server in a goroutine
	go func() {
		main()
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	t.Run("TestRootEndpoint", func(t *testing.T) {
		testRequest(t, "GET / HTTP/1.1\r\n\r\n", "HTTP/1.1 200 OK", "")
	})

	t.Run("TestNotFoundEndpoint", func(t *testing.T) {
		testRequest(t, "GET /abcdefg HTTP/1.1\r\n\r\n", "HTTP/1.1 404 Not Found", "")
	})

	t.Run("TestEchoEndpoint", func(t *testing.T) {
		testRequest(t, "GET /echo/abc HTTP/1.1\r\n\r\n", "HTTP/1.1 200 OK", "abc")
	})

	t.Run("TestUserAgentEndpoint", func(t *testing.T) {
		testRequest(t, "GET /user-agent HTTP/1.1\r\nUser-Agent: foobar/1.2.3\r\n\r\n",
			"HTTP/1.1 200 OK", "foobar/1.2.3")
	})

	t.Run("TestConcurrentConnections", func(t *testing.T) {
		numRequests := 4
		done := make(chan bool, numRequests)

		for i := 0; i < numRequests; i++ {
			go func() {
				testRequest(t, "GET / HTTP/1.1\r\n\r\n", "HTTP/1.1 200 OK", "")
				done <- true
			}()
		}

		// Wait for all requests to complete
		for i := 0; i < numRequests; i++ {
			<-done
		}
	})
}

func testRequest(t *testing.T, request string, expectedCode string, expectedBody string) {
	conn, err := net.Dial("tcp", "localhost:4221")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("Failed to write to connection: %v", err)
	}

	reader := bufio.NewReader(conn)
	responseLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if !strings.HasPrefix(responseLine, expectedCode) {
		t.Errorf("Expected response to start with %q, but got %q", expectedCode, responseLine)
	}

	// Skip the rest if expected body is empty
	if expectedBody == "" {
		return
	}

	// Store the last non-empty line
	var body string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("BREAKING")
				break
			}
		}
		line = strings.TrimSpace(line)

		if line != "" {
			body = line
		}
	}

	if expectedBody != body {
		t.Errorf("Expected body to be %q, but got %q", expectedBody, body)
	}
}
