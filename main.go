package main

import (
	"log"
	"net"
)


func main() {
	chatServer := CreateServer()
	go chatServer.run()
	addressName := ":8080"
	listener, err := net.Listen("tcp", addressName)
	if err != nil {
		log.Fatalf("Unable to start server %s", addressName)
	}
	defer listener.Close()
	log.Printf("Starting server at %s", addressName)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error in listening a connection: %s", err.Error())
			continue
		}
		c := chatServer.newClient(conn)
		go c.readInput()
	}
}