package serv

import (
	"bytes"
	"net"
	"strings"
	"testing"
	"time"
)

// MockConn is a mock implementation of net.Conn for testing purposes.
type MockConn struct {
	ReadBuffer  *bytes.Buffer
	WriteBuffer *bytes.Buffer
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.ReadBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.WriteBuffer.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return nil
}

func (m *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// TestHandleConnection tests the handleConnection function.
func TestHandleConnection(t *testing.T) {
	// Prepare a mock connection with a name.
	mockConn := &MockConn{
		ReadBuffer:  bytes.NewBufferString("Alice\n"),
		WriteBuffer: &bytes.Buffer{},
	}

	// Call handleConnection in a goroutine.
	go handleConnection(mockConn)

	// Wait for a moment to allow the goroutine to process.
	time.Sleep(100 * time.Millisecond)

	// Check if the welcome message was sent.
	if !strings.Contains(mockConn.WriteBuffer.String(), "Welcome to TCP-Chat!") {
		t.Errorf("Expected welcome message, got: %s", mockConn.WriteBuffer.String())
	}

	// Simulate sending a message from Alice.
	mockConn.ReadBuffer = bytes.NewBufferString("Hello!\n")
	go handleConnection(mockConn)

	// Wait for a moment to allow broadcasting.
	time.Sleep(100 * time.Millisecond)

	// Check if the message was broadcasted correctly.
	if !strings.Contains(mockConn.WriteBuffer.String(), "Hello!") {
		t.Errorf("Expected broadcast message, got: %s", mockConn.WriteBuffer.String())
	}
}

// TestBroadcast tests the broadcast function.
func TestBroadcast(t *testing.T) {
	mockConn1 := &MockConn{
		WriteBuffer: &bytes.Buffer{},
	}
	mockConn2 := &MockConn{
		WriteBuffer: &bytes.Buffer{},
	}

	mu.Lock()
	clients[mockConn1] = "Alice"
	clients[mockConn2] = "Bob"
	mu.Unlock()

	broadcast("Test message")

	time.Sleep(100 * time.Millisecond)

	if !strings.Contains(mockConn1.WriteBuffer.String(), "Test message") {
		t.Error("Client 1 did not receive broadcasted message.")
	}
	if !strings.Contains(mockConn2.WriteBuffer.String(), "Test message") {
		t.Error("Client 2 did not receive broadcasted message.")
	}
}

// TestEmptyName checks that an empty name is handled correctly.
func TestEmptyName(t *testing.T) {
	mockConn := &MockConn{
		ReadBuffer:  bytes.NewBufferString("\n"),
		WriteBuffer: &bytes.Buffer{},
	}

	handleConnection(mockConn)

	time.Sleep(100 * time.Millisecond)

	if !strings.Contains(mockConn.WriteBuffer.String(), "Name cannot be empty!") {
		t.Error("Expected 'Name cannot be empty!' message.")
	}
}

// Cleanup function to reset clients after each test
func cleanup() {
	mu.Lock()
	defer mu.Unlock()
	for conn := range clients {
		delete(clients, conn)
	}
}

func TestMain(m *testing.M) {
	cleanup()
	m.Run()
	cleanup()
}
