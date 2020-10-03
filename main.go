package main

import (
	"fmt"

	"example.com/rickgrammer/dsa/ll/dll"
	"example.com/rickgrammer/dsa/ll/sll"
)

func main() {
	keys := []int{1,2,3,4}
	d := dll.Create()
	d.BulkAdd(keys...)
	d.Add(3)
	d.Add(4)
	d.Add(6)
	d.Delete(6)
	d.Display()
	fmt.Println(d.Search(8))
	fmt.Println(d.Search(3))

	l := sll.Create()
	// l.Add(1)
	l.BulkAdd(keys...)
	l.Add(2)
	l.Add(3)
	l.Display()
	// l.Delete(2)
	// fmt.Println(l.Search(1))
	// l.Display()
	l.Reverse()
	l.Display()

}
