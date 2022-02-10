package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type payload struct {
	TYPE string `json:"type"`
	EVENT string `json:"event"`
	DATA string `json:"data"`
}

var workers []*net.TCPConn

func main() {
	stream, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	defer stream.Close()

	for {
		conn, err := stream.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			os.Exit(1)
		}
		fmt.Println("got a connection")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var haveDeclared = false
	logger := log.New(os.Stdout, "", 0)
	for {
		d := json.NewDecoder(conn)
		var msg payload

		d.Decode(&msg)

		if !haveDeclared {
			logger.Println("Declaring")
			if msg.TYPE == "worker" {
				tcp := conn.(*net.TCPConn)
				workers = append(workers, tcp)
			}
			haveDeclared = true
		} else {
			logger.Println("In else")
			for _, worker := range workers {
				logger.Println("looping")
				b, err := json.Marshal(msg)
				if err != nil {
					continue
				}
				logger.Println("Writing")
				worker.Write(b)
			}
		}
	}
}