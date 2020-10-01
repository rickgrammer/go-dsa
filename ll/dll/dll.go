package dll

import (
	"fmt"
)

type node struct {
	prev *node
	key  int
	next *node
}

type Dll struct {
	sentinel *node
}

func (d *Dll) Display() {
	for current := d.sentinel.next; current != d.sentinel; current = current.next {
		fmt.Print(current.key)
		if (current.next != d.sentinel){
			fmt.Print("-")
		}
	}
	fmt.Println()
}

func (d *Dll) search(key int) *node {
	for current := d.sentinel.next; current != d.sentinel; current = current.next {
		if key == current.key {
			return current
		}
	}
	return nil
}

func (d *Dll) Search(key int) bool {
	return !(d.search(key) == nil)
}

func (d *Dll) Delete(key int) {
	found := d.search(key)
	if found == nil {
		return
	}
	if found.prev != nil {
		found.prev.next = found.next
	}
	if found.next != nil {
		found.next.prev = found.prev
	}
}

// Create a Doubly-Linked List with sentinel
func Create() *Dll {
	sentinal := node{}
	sentinal.next, sentinal.prev = &sentinal, &sentinal
	return &Dll{
		sentinel: &sentinal,
	}
}

func (d *Dll) Add(key int) {
	d.sentinel.next = &node{
		prev: d.sentinel,
		key: key,
		next: d.sentinel.next,
	}
}
