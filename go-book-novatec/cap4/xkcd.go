package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Movie struct {
	Num                                                                   int
	Day, Year, Month, Link, News, Safe_title, Transcript, Alt, Img, Title string
}

const urlStart = "https://xkcd.com/"
const urlEnd = "/info.0.json"

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(urlStart + url + urlEnd)
		if err != nil {
			println("Algo deu errado")
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println(resp.Status)
		}

		var result = &Movie{}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			fmt.Println("er ", resp.Status, err)
		}

		fmt.Println(*result)
	}
}
