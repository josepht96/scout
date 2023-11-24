package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func dial(conn *net.UDPConn) {
	// Send a message to the server
	preTime := time.Now()
	preTimeStr := preTime.Format(time.RFC3339Nano)
	log.Printf("Start time: %s", preTimeStr)
	_, err := conn.Write([]byte(preTimeStr))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is sent
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	postTime := time.Now()
	rtt := postTime.Sub(preTime).String()

	dataStr := strings.TrimSuffix(data, "\n")
	serverReceivedTime, err := time.Parse(time.RFC3339Nano, dataStr)
	if err != nil {
		log.Fatalln(err)
	}
	latency := serverReceivedTime.Sub(preTime).String()

	log.Printf("Latency to server: %s", string(latency))
	log.Printf("RTT: %s", string(rtt))
}
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port to connect to")
		os.Exit(1)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%s", os.Args[1]))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(udpAddr)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {
		dial(conn)
		time.Sleep(3 * time.Second)
	}

}
