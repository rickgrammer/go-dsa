package main

import (
	"fmt"
	"example.com/rickgrammer/dsa/ll/dll"
	"example.com/rickgrammer/dsa/ll/sll"
)

func main() {
	d := dll.Create()
	d.Add(3)
	d.Add(4)
	d.Add(6)
	d.Delete(6)
	d.Display()
	fmt.Println(d.Search(8))
	fmt.Println(d.Search(3))

	l := sll.Create()
	l.Add(4)
	l.Add(2)
	l.Add(1)
	l.Add(4)
	l.Add(2)
	l.Add(1)
	l.Display()
	l.Delete(2)
	fmt.Println(l.Search(1))
	l.Display()

}
