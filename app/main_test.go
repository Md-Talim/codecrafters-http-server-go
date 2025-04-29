package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

// TestMain runs setup before all tests in the package and teardown after.
func TestMain(m *testing.M) {
	// Setup: Start the server in a goroutine
	fmt.Println("Starting server for tests...")
	go main()
	// Give the server time to start listening.
	time.Sleep(500 * time.Millisecond)

	// Run all tests in the package
	exitCode := m.Run()

	fmt.Println("Finished running tests.")
	os.Exit(exitCode)
}

func TestServerBindsToPort(t *testing.T) {
	conn, err := net.DialTimeout("tcp", "localhost:4221", 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
}

func TestResponsdWith200OK(t *testing.T) {
	// 1. Connect to the server
	conn, err := net.DialTimeout("tcp", "localhost:4221", 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 2. Send a minimal HTTP GET request
	request := "GET / HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: test-client\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// 3. Read the response
	// Create a buffer to read the response into. 1024 bytes should be plenty.
	buffer := make([]byte, 1024)
	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		t.Fatalf("Failed to set read deadline: %v", err)
	}

	n, err := conn.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	// 4. Assert the response content
	expectedResponse := "HTTP/1.1 200 OK\r\n\r\n"
	actualResponse := buffer[:n] // Get the slice containing the actual data read

	if !bytes.Equal(actualResponse, []byte(expectedResponse)) {
		t.Errorf("Unexpected response:\nExpected:\n%q\nGot:\n%q", expectedResponse, actualResponse)
	}
}

func TestRespondWith404_Not_Found(t *testing.T) {
	conn, err := net.DialTimeout("tcp", "localhost:4221", 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	requestPath := "/some-random-path"
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: test-client/0.0.1\r\nAccept: */*\r\n\r\n", requestPath)
	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	buffer := make([]byte, 1024)                                // Use a reasonably sized buffer
	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // Set a deadline
	if err != nil {
		t.Fatalf("Failed to set read deadline: %v", err)
	}

	n, err := conn.Read(buffer)
	if err != nil {
		// Check if it's a timeout error
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			t.Fatalf("Timeout reading response: Server might not have responded in time. %v", err)
		}
		t.Fatalf("Failed to read response: %v", err)
	}
	if n == 0 {
		t.Fatal("Read 0 bytes from connection, expected a response.")
	}

	expectedResponse := "HTTP/1.1 404 Not Found\r\n\r\nNot Found"
	actualResponse := buffer[:n] // Get the slice containing the actual data read

	if !bytes.Equal(actualResponse, []byte(expectedResponse)) {
		t.Errorf("Unexpected response for GET %s:\nExpected:\n%q\nGot:\n%q", requestPath, expectedResponse, actualResponse)
	}
}

func TestRespondWithBody(t *testing.T) {
	conn, err := net.DialTimeout("tcp", "localhost:4221", 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	requestPath := "/echo/abc"
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: test-client/0.0.1\r\nAccept: */*\r\n\r\n", requestPath)
	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	buffer := make([]byte, 1024)                                // Use a reasonably sized buffer
	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // Set a deadline
	if err != nil {
		t.Fatalf("Failed to set read deadline: %v", err)
	}

	n, err := conn.Read(buffer)
	if err != nil {
		// Check if it's a timeout error
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			t.Fatalf("Timeout reading response: Server might not have responded in time. %v", err)
		}
		t.Fatalf("Failed to read response: %v", err)
	}
	if n == 0 {
		t.Fatal("Read 0 bytes from connection, expected a response.")
	}

	expectedResponse := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nabc"
	actualResponse := buffer[:n] // Get the slice containing the actual data read

	if !bytes.Equal(actualResponse, []byte(expectedResponse)) {
		t.Errorf("Unexpected response for GET %s:\nExpected:\n%q\nGot:\n%q", requestPath, expectedResponse, actualResponse)
	}
}
