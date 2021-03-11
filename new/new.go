package main

import "fmt"

type P struct{
	Name string
	Age int
}

func main() {
	var ps *P
	ps = new(P)

	fmt.Println(ps)
	fmt.Println(*ps)
}