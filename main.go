package main

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/divy-sh/animus/command"
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func main() {
	Handle()
}

// retry executes a function up to maxRetries times with a delay between attempts
func retry(maxRetries int, delay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		log.Printf("Attempt %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(delay)
	}
	return err
}

func Handle() {
	log.Print("Listening on port :6379")

	var l net.Listener

	// Retry listener creation
	err := retry(5, 2*time.Second, func() error {
		var err error
		l, err = net.Listen("tcp", ":6379")
		return err
	})

	if err != nil {
		log.Printf("Failed to start server after retries: %v", err)
		return
	}
	defer l.Close()
	log.Print("Server started successfully")

	for {
		var conn net.Conn

		// Retry accepting connection
		err := retry(5, 2*time.Second, func() error {
			var err error
			conn, err = l.Accept()
			return err
		})

		if err != nil {
			log.Printf("Failed to accept connection after retries: %v", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			handleRequests(c)
		}(conn)
	}
}

func handleRequests(conn net.Conn) {
	reader := resp.NewReader(conn)
	writer := resp.NewWriter(conn)
	for {
		value, err := reader.Read()
		if err != nil {
			// log.Print("Error reading request:", err)
			return
		}
		if value.Typ != "array" || len(value.Array) == 0 {
			log.Print("Invalid request, expected array")
			writer.Write(resp.Value{Typ: common.STRING_TYPE, Str: "Invalid request"})
			continue
		}
		cmd := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]
		if cmd == "QUIT" {
			writer.Write(resp.Value{Typ: common.STRING_TYPE, Str: "OK"})
			return
		}
		handler, ok := command.Handlers[cmd]
		if !ok {
			// log.Print("Invalid command: ", cmd)
			writer.Write(resp.Value{Typ: common.STRING_TYPE, Str: "Invalid command"})
			continue
		}
		result := handler.Func(args)
		writer.Write(result)
	}
}
