package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

var address = flag.String("address", ":15000", "Address to listen on.")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("echo server listening on %s", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(conn, "macOS service received: %s\n", scanner.Text()); err != nil {
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("connection failed: %v", err)
	}
}
