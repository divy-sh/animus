package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("listening on port:6379")
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			os.Exit(1)
		}
		conn.Write([]byte("+OK\r\n"))
	}
}
