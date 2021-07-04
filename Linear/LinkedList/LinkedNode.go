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

func NewLinkedNodeSimplified(value interface{}, next *LinkedNodeSimplified) *LinkedNodeSimplified {
	return &LinkedNodeSimplified{
		Value: value,
		Next:  next,
	}
}

func NewLinkedNodeSimplifiedByValue(value interface{}) *LinkedNodeSimplified {
	return NewLinkedNodeSimplified(value, nil)
}

func NewLinkedNode(value interface{}, next, prev *LinkedNode) *LinkedNode {
	return &LinkedNode{
		Value: value,
		Next:  next,
		Prev:  prev,
	}
}

func NewLinkedNodeByValue(value interface{}) *LinkedNode {
	return NewLinkedNode(value, nil, nil)
}
