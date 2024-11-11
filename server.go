package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func sendResponse(conn net.Conn, code int, msg string, body string) {
	headers := fmt.Sprintf("HTTP/1.1 %d %s\r\n\r\n", code, msg)
	headers += "Content-Type: text/plain\r\n"
	headers += fmt.Sprintf("Content-Length: %d \r\n\r\n%s", len(body), body)
	headers += "\r\n"

	conn.Write([]byte(headers))
}

func routeRequest(path string, conn net.Conn) {
	switch {
	case path == "/":
		sendResponse(conn, 200, "OK", "")
	case strings.HasPrefix(path, "/echo"):
		rbody := strings.TrimPrefix(path, "/echo/")
		sendResponse(conn, 200, "OK", rbody)
	default:
		sendResponse(conn, 404, "Not Found", "")
	}
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

	routeRequest(path, conn)
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
