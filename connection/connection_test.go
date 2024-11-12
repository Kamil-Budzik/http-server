package connection

import (
	"bytes"
	"net"
	"testing"
	"time"
)

// MockConn is a mock implementation of the net.Conn interface for testing.
type MockConn struct {
	readBuffer  bytes.Buffer
	writeBuffer bytes.Buffer
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.readBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.writeBuffer.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr                { return nil }
func (m *MockConn) RemoteAddr() net.Addr               { return nil }
func (m *MockConn) SetDeadline(t time.Time) error      { return nil }
func (m *MockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockConn) SetWriteDeadline(t time.Time) error { return nil }

// TestSendResponse tests the SendResponse method of the Connection struct.
// func TestSendResponse(t *testing.T) {
// 	mockConn := &MockConn{}
// 	connection := NewConnection(mockConn)
//
// 	code := 200
// 	msg := "OK"
// 	body := "Hello, World!"
//
// 	connection.SendResponse(code, msg, body)
//
// 	expectedOutput := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nHello, World!\r/n"
// 	if mockConn.writeBuffer.String() != expectedOutput {
// 		t.Errorf("expected %q, got %q", expectedOutput, mockConn.writeBuffer.String())
// 	}
// }

// TestGetHeaderValue tests the GetHeaderValue method of the Connection struct.
func TestGetHeaderValue(t *testing.T) {
	mockConn := &MockConn{}
	mockConn.readBuffer.WriteString("Content-Type: text/plain\r\nContent-Length: 13\r\n\r\n")

	connection := NewConnection(mockConn)

	headerValue := connection.GetHeaderValue("Content-Length")
	expectedValue := "13"

	if headerValue != expectedValue {
		t.Errorf("expected %q, got %q", expectedValue, headerValue)
	}
}
