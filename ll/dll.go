package ll

type Node struct {
	Prev *Node
	Key int
	Next *Node
}


func Create(key int) *Node {
	return &Node{
		Key: key,
	}
}