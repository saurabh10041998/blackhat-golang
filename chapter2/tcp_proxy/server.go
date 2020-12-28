package main

import (
	"io"
	"log"
	"net"
)

// Echo is just function which echos back msg
func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

func main() {
	// starting a server at 3000
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalln("Unable to bind port 3000")
	}
	log.Println("Server started listening at 0.0.0.0:3000")

	// Wait for the connection and create conn for every connection received
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln("Unable to receive connection from client")
		}
		// handle each connection by creating the go routine for each request
		go echo(conn)
	}
}
