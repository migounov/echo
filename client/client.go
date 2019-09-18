package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const addr = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(conn)

	for {
		fmt.Print("Enter message: ")
		message, err := reader.ReadBytes('\n')
		switch err {
		case nil:
			_, _ = conn.Write(message)
		case io.EOF:
			os.Exit(0)
		default:
			log.Fatal(err)
		}

		r, err := response.ReadBytes(byte('\n'))

		switch err {
		case nil:
			fmt.Print(string(r))
		case io.EOF:
			os.Exit(0)
		default:
			log.Fatal(err)
		}
	}
}
