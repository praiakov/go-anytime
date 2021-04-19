// Fetchall busca URLs em paralelo e informe os tempos gastos e os tamanhos
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma gorroutina
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // recebe do canal ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //envia para o canal ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	f, err := os.Create("../fetchall.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	f2, err := f.Write([]byte(strconv.FormatInt(nbytes, 10)))

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(f2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	f.Write([]byte(" " + url))
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
