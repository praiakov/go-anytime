package main

// waitforserver tenta contatar o servidor de uma URL

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// tenta por um minuto usando exponential back-off
// devolve um erro se todas as tentativas falharem
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // sucesso
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := "https://praia.dev/"
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}
