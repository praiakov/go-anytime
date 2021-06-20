package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // um canal de mensagem de saída

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // todas as mensagens que chegam dos clientes
)

func broadcaster() {
	clients := make(map[client]bool) //todos os clients conectados
	for {
		select {
		case msg := <-messages:
			// Broadcast de mensagem de entrada para todos
			// os canais de mensagem ed saída dos clientes
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // mensagens de saída do cliente
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTA: ignorando erros em potencial de input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTA: ignorando erros de rede
	}
}
