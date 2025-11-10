package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/divy-sh/animus/resp"
)

func TestHandleRequests_Ping(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	// Run the server side
	go handleRequests(server)

	writer := resp.NewWriter(client)
	reader := resp.NewReader(client)

	// Send a RESP array: ["PING"]
	writer.Write(resp.Value{
		Typ: "array",
		Array: []resp.Value{
			{Typ: "bulk", Bulk: "PING"},
		},
	})

	// Read response
	value, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if value.Array[0].Bulk != "+PONG" {
		t.Fatalf("Expected PONG, got %q", value.Str)
	}
}

func TestHandleRequests_Quit(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go handleRequests(server)

	writer := resp.NewWriter(client)
	reader := resp.NewReader(client)

	// Send QUIT command
	writer.Write(resp.Value{
		Typ: "array",
		Array: []resp.Value{
			{Typ: "bulk", Bulk: "QUIT"},
		},
	})

	value, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if value.Array[0].Bulk != "+OK" {
		t.Fatalf("Expected OK, got %q", value.Str)
	}
}

func Test_HandleRequests_InvalidCommand(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go handleRequests(server)

	writer := resp.NewWriter(client)
	reader := resp.NewReader(client)

	// Send an invalid command
	writer.Write(resp.Value{
		Typ: "array",
		Array: []resp.Value{
			{Typ: "bulk", Bulk: "INVALIDCMD"},
		},
	})

	value, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if value.Array[0].Bulk != "+Invalid" || value.Array[1].Bulk != "command" {
		t.Fatalf("Expected: +Invalid command, got %q", value)
	}
}

func Test_HandleRequests_EmptyCommand(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go handleRequests(server)

	writer := resp.NewWriter(client)
	reader := resp.NewReader(client)

	// Send an invalid command
	writer.Write(resp.Value{
		Typ:   "array",
		Array: []resp.Value{},
	})

	value, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if value.Array[0].Bulk != "+Invalid" || value.Array[1].Bulk != "request" {
		t.Fatalf("Expected: +Invalid command, got %q", value)
	}
}

func TestHandle(t *testing.T) {
	// Run the server in a goroutine
	go main()
	if err := waitForServer("127.0.0.1:6379", time.Second); err != nil {
		t.Fatalf("server never started: %v", err)
	}

	// Try connecting to the hardcoded port
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	writer := resp.NewWriter(conn)
	reader := resp.NewReader(conn)

	// Send a PING command
	writer.Write(resp.Value{
		Typ: "array",
		Array: []resp.Value{
			{Typ: "bulk", Bulk: "PING"},
		},
	})

	// Read the response
	respVal, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if respVal.Array[0].Bulk != "+PONG" {
		t.Errorf("Expected PONG, got %v", respVal)
	}

	// Cleanly close server connection by sending QUIT
	writer.Write(resp.Value{
		Typ: "array",
		Array: []resp.Value{
			{Typ: "bulk", Bulk: "QUIT"},
		},
	})

	// Optional: wait a bit for server goroutine to finish
	time.Sleep(100 * time.Millisecond)
}

func waitForServer(addr string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("server did not start on time")
}
