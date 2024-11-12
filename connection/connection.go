package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Connection struct {
	Conn   net.Conn
	Reader *bufio.Reader
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn:   conn,
		Reader: bufio.NewReader(conn),
	}
}

func (c *Connection) SendResponse(code int, msg string, contentType string, body string) {
	headers := fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, msg)
	headers += fmt.Sprintf("Content-Type: %s\r\n", contentType)
	headers += fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	headers += body
	headers += "\r\n"

	c.Conn.Write([]byte(headers))
}

func (c *Connection) GetHeaderValue(header string) string {
	for {
		line, err := c.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading header:", err)
			return ""
		}

		if strings.HasPrefix(line, header) {
			segments := strings.Split(line, ": ")
			return strings.TrimSuffix(segments[len(segments)-1], "\r\n")

		}
	}
}
