package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	socket, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	defer socket.Close()

	conn, connErr := socket.Accept()
	if connErr != nil {
		log.Fatal(connErr)
	}

	for {
		clientData, clientErr := bufio.NewReader(conn).ReadString('\n')
		if clientErr != nil {
			log.Fatal(clientErr)
		}
		if strings.TrimSpace(string(clientData)) == "STOP" {
			conn.Write([]byte("Exiting!!!"))
			log.Println("Exiting TCP server!!!!!")
			return
		}
		log.Println("-> ", string(clientData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))
	}

}
