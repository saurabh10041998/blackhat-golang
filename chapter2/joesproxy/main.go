package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		log.Fatalln("unable to connect to required host")
	}

	defer dst.Close()

	// go routine to prevent the blocking of io.Copy
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// let keep this running
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	//listen on local port
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("unable to bind port 80")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
