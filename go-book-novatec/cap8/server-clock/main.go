package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // por ex, conexao abortadada
		}
		handleCon(conn) // trata uma conexao de cada vez
	}
}

func handleCon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // por ex, cliente desconectou
		}
		time.Sleep(1 * time.Second)
	}
}
