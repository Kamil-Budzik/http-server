package main

import (
	"flag"
	"fmt"
	"http-server/connection"
	"log"
	"net"
	"strings"
)

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

	// Handle routing
	switch {
	case path == "/":
		clientConn.SendResponse(200, "OK", "")
	case strings.HasPrefix(path, "/echo"):
		responseBody := strings.TrimPrefix(path, "/echo/")
		clientConn.SendResponse(200, "OK", responseBody)
	case path == "/user-agent":
		userHeader := clientConn.GetHeaderValue("User-Agent")
		clientConn.SendResponse(200, "OK", userHeader)
	default:
		clientConn.SendResponse(404, "Not Found", "")
	}
}

type Server struct {
	Address       string
	Port          int
	ActiveClients int
	ShutdownChan  chan struct{}
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}
}

var directory string

func main() {
	// Read directory flag
	flag.StringVar(&directory, "directory", "tmp", "specifies the directory where the files are stored, as an absolute path")
	flag.Parse()
	fmt.Println("DIRECTORY", directory)

	server := &Server{
		Address:      "localhost",
		Port:         4221,
		ShutdownChan: make(chan struct{}),
	}
	server.start()
	fmt.Println("Server started at port 4221")

}
