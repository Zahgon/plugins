// Echosvr is a simple TCP echo server
//
// It prints its listen address on stdout
//
//	  127.0.0.1:xxxxx
//	A test should wait for this line, parse it
//	and may then attempt to connect.
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Start TCP server
	listener, err := net.Listen("tcp", ":")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	// use the same port for UDP
	_, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("127.0.0.1:%s\n", port)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			go handleConnection(conn)
		}
	}()

	// Start UDP server
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("Error from net.ResolveUDPAddr(): %s", err)
		return
	}
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Printf("Error from ListenUDP(): %s", err)
		return
	}
	defer sock.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := sock.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error from ReadFrom(): %s", err)
			return
		}
		sock.SetWriteDeadline(time.Now().Add(1 * time.Minute))
		_, err = sock.WriteTo(buffer[0:n], addr)
		if err != nil {
			return
		}
	}
}

func handleConnection(conn net.Conn) { _ = "STUB: not implemented"; return }
