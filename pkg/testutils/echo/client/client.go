package main

import (
	"flag"
)

func main() {
	target := flag.String("target", "", "the server address")
	payload := flag.String("message", "", "the message to send to the server")
	protocol := flag.String("protocol", "tcp", "the protocol to use with the server [udp,tcp], default tcp")
	flag.Parse()

	if *target == "" || *payload == "" {
		flag.Usage()
		panic("invalid arguments")
	}

	switch *protocol {
	case "tcp":
		connectTCP(*target, *payload)
	case "udp":
		connectUDP(*target, *payload)
	default:
		panic("invalid protocol")
	}
}

func connectTCP(target, payload string) { _ = "STUB: not implemented"; return }

// UDP uses a constant source port to trigger conntrack problems
func connectUDP(target, payload string) { _ = "STUB: not implemented"; return }
