package main

import "fmt"

type strategy2 struct {
}

func (l *strategy2) process() {
	fmt.Println("strategy 2")
}
