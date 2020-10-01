package sll

import (
	"fmt"
)

type node struct {
	key  int
	next *node
}

type Sll struct {
	head *node
}

func (s *Sll) Display() {
	for current := s.head.next; current != nil; current = current.next {
		fmt.Print(current.key)
		if current.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (s *Sll) search(key int) *node {
	for current := s.head.next; current != nil; current = current.next {
		if key == current.key {
			return current
		}
	}
	return nil
}

func (s *Sll) Search(key int) bool {
	return !(s.search(key) == nil)
}

func (s *Sll) delete(prev, current *node) {
	prev.next = current.next
}

func (s *Sll) Delete(key int) {
	var prev *node
	for current := s.head.next; current != nil; current = current.next {
		if (key == current.key) {
			s.delete(prev, current)	
			return
		}
		prev = current
	}
	return
}

func (s *Sll) Add(key int) {
	s.head.next = &node{
		key:  key,
		next: s.head.next,
	}
}

func Create() *Sll {
	return &Sll{
		head: &node{},
	}
}
