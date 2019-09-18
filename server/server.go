package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
)

const port = "localhost:8080"

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			log.Fatal(err)
		}
		_, _ = conn.Write(bytes.ToUpper(line))
	}
}

func main() {
	log.Println("Starting server, listening on port " + port)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go echo(conn)
	}
}