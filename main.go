package main

import (
	"log"
	"net"
	"strings"

	"github.com/divy-sh/animus/internal/command"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func main() {
	log.Print("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Print(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		log.Print(err)
		return
	}

	defer conn.Close()

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

		cmd := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)

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
