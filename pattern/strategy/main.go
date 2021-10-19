package main

func main() {
	m := merchant{&strategy1{}}
	m.Messager.process()

	a := merchant{&strategy2{}}
	a.Messager.process()
}
