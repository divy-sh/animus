package main

import (
	"log"
	"net"
	"strings"

	"github.com/divy-sh/animus/commands"
	"github.com/divy-sh/animus/resp"
)

func main() {
	log.Print("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Print(err)
		return
	}

	aof, err := NewAof("database.aof")
	if err != nil {
		log.Print(err)
		return
	}
	defer aof.Close()

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		log.Print(err)
		return
	}

	defer conn.Close()

	aof.Read(func(value resp.Value) {
		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := commands.Handlers[command]
		if !ok {
			log.Print("Invalid command: ", command)
		}
		handler.Func(args)
	})

	for {
		reader := resp.NewReader(conn)
		value, err := reader.Read()
		if err != nil {
			log.Print(err)
			return
		}

		if value.Typ != "array" {
			log.Print("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			log.Print("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)

		handler, ok := commands.Handlers[command]
		if !ok {
			log.Print("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}

		result := handler.Func(args)
		writer.Write(result)
	}
}
