package main

import (
	"net"
	"testing"
	"time"
)

func TestServerBindsToPort(t *testing.T) {
	go main() // Start the server in a goroutine

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "localhost:4221")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
}
