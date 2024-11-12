package main

import (
	"bufio"
	"flag"
	"fmt"
	"http-server/connection"
	"log"
	"net"
	"os"
	"strings"
)

func ReadFileLines(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (s *Server) handleConnection(conn net.Conn) {
	s.ActiveClients++
	defer func() {
		conn.Close()
		s.ActiveClients--
	}()

	clientConn := connection.NewConnection(conn)

	requestLine, err := clientConn.Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")

	if len(parts) != 3 {
		fmt.Println("Invalid request line:", requestLine)
		return
	}

	method := parts[0]
	path := parts[1]
	version := parts[2]

	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Version: %s\n", version)

	switch {
	case path == "/":
		clientConn.SendResponse(200, "OK", "text/plain", "")
	case strings.HasPrefix(path, "/echo"):
		responseBody := strings.TrimPrefix(path, "/echo/")
		clientConn.SendResponse(200, "OK", "text/plain", responseBody)
	case path == "/user-agent":
		userHeader := clientConn.GetHeaderValue("User-Agent")
		clientConn.SendResponse(200, "OK", "text/plain", userHeader)
	case strings.HasPrefix(path, "/files"):
		fileName := strings.TrimPrefix(path, "/files/")
		fileName = fmt.Sprintf("./%s/%s", s.Directory, fileName)
		contents, err := ReadFileLines(fileName)

		if err != nil {
			clientConn.SendResponse(404, "Not Found", "application/octet-stream", "")
		}

		clientConn.SendResponse(200, "OK", "application/octet-stream", contents)
	default:
		clientConn.SendResponse(404, "Not Found", "text/plain", "")
	}
}

type Server struct {
	Address       string
	Port          int
	ActiveClients int
	ShutdownChan  chan struct{}
	Directory     string
}

func (s *Server) start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Address, s.Port))
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	go func() {
		<-s.ShutdownChan
		fmt.Println("Server shutting down...")
		listener.Close()
	}()

	fmt.Printf("Server started at port %d \n", s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func main() {
	server := &Server{
		Address:      "localhost",
		Port:         4221,
		ShutdownChan: make(chan struct{}),
	}

	flag.StringVar(&server.Directory, "directory", "tmp", "specifies the directory where the files are stored, as an absolute path")
	flag.Parse()

	server.start()
}
