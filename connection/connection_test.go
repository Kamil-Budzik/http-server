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
func TestSendResponse(t *testing.T) {
	mockConn := &MockConn{}
	connection := NewConnection(mockConn)

	code := 200
	msg := "OK"
	contentType := "text/plain"
	body := "Hello, World!"

	connection.SendResponse(code, msg, contentType, body)

	expectedOutput := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nHello, World!\r\n"
	if mockConn.writeBuffer.String() != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, mockConn.writeBuffer.String())
	}
}

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

func TestReadBody(t *testing.T) {
	headers := "POST /test HTTP/1.1\r\nContent-Length: 5\r\n\r\n"
	body := "12345"
	rawRequest := headers + body

	mockConn := &MockConn{}
	mockConn.readBuffer.WriteString(rawRequest)

	connection := NewConnection(mockConn)

	result := connection.ReadBody()

	expected := "12345"
	if result != expected {
		t.Errorf("expected body %q, got %q", expected, result)
	}
}
