package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%s", os.Args[1]))
	if err != nil {
		log.Fatalln(err)
	}

	// Start listening for UDP packages on the given address
	log.Printf("Starting UDP server at %s", fmt.Sprintf("localhost:%s", os.Args[1]))
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Read from UDP listener in endless loop
	for {
		var buf [256]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		now := time.Now()
		nowStr := now.Format(time.RFC3339Nano)
		// reqBytes := bytes.Trim(buf[:], "\x00")
		// sentTime, err := time.Parse(time.RFC3339Nano, string(reqBytes))
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// latency := now.Sub(sentTime).String()
		// log.Printf("Latency to server: %s", string(latency))

		conn.WriteToUDP([]byte(fmt.Sprintf("%s\n", nowStr)), addr)
	}
}
