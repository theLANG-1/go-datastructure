package LinkedList

type LinkedNodeSimplified struct {
	Value interface{}
	Next  *LinkedNodeSimplified
}

type LinkedNode struct {
	Value interface{}
	Next  *LinkedNode
	Prev  *LinkedNode
}
