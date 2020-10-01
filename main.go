package main

import (
	"fmt"

	"example.com/rickgrammer/dsa/ll"
)

func main() {
	fmt.Println("HI")
	x := ll.Node{
		Key: 2,
	}
	fmt.Printf("%p\n", &x)
	x.Key = 4
	fmt.Printf("%p\n", &x)
}
