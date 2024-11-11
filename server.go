package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func sendResponse(conn net.Conn, code int, msg string) {
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n\r\n", code, msg)
	conn.Write([]byte(response))
}

func extractPathParameters(path string) string {
	segments := strings.Split(path, "/")
	fmt.Println(segments)
	if len(segments) > 1 {
		return segments[len(segments)-1]
	}
	return ""
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")

	if len(parts) != 3 {
		fmt.Println("Invalid request line: ", requestLine)
		return
	}

	method := parts[0]
	path := parts[1]
	version := parts[2]

	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Version: %s\n", version)

	switch path {
	case "/":
		sendResponse(conn, 200, "OK")
	case "/echo":
		rbody := extractPathParameters(path)
		sendResponse(conn, 200, rbody)
	default:
		sendResponse(conn, 404, "Not Found")
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer listen.Close()
	fmt.Println("Server started at port 4221")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error acceptng connection:", err)
			continue
		}

		go handleConnection(conn)
	}

}
