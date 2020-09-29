package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		log.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(socket, text+"\n")

		message, _ := bufio.NewReader(socket).ReadString('\n')
		log.Println("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			log.Println("TCP client exiting.....")
			return
		}
	}
}
