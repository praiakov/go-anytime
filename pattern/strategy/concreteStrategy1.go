package main

import "fmt"

type strategy1 struct {
}

func (l *strategy1) process() {

	fmt.Println("strategy 1")

}
