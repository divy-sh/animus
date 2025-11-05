package main

import (
	"log"
	"net"
	"strings"

	"github.com/divy-sh/animus/command"
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func main() {
	Handle()
}

func Handle() {
	log.Print("Listening on port :6379")
	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Print(err)
		return
	}
	for {
		// Listen for connections
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		// Handle each connection in its own goroutine and ensure it is closed
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
			log.Print("Error reading request:", err)
			return
		}
		if value.Typ != "array" || len(value.Array) == 0 {
			log.Print("Invalid request, expected array")
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
			log.Print("Invalid command: ", cmd)
			writer.Write(resp.Value{Typ: common.STRING_TYPE, Str: ""})
			continue
		}
		result := handler.Func(args)
		writer.Write(result)
	}
}
