package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	_ = requestLine

	conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\n"))
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	_, err = listen.Accept()
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
