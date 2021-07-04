package LinkedList

import (
	"fmt"
	"strings"
)

// LinkedList use LinkedNode as its node, and LinkedNode
// has the pointer named Next and Prev which make it more
// easily for LinkedList to find or remove a node.

type linkedList struct {
	virtualHead *LinkedNode
	last        *LinkedNode
	len         int
}

func (l *linkedList) AddFirst(value interface{}) {
	var (
		tmpNode = NewLinkedNode(value, l.virtualHead.Next, l.virtualHead)
	)
	if l.virtualHead.Next != nil {
		l.virtualHead.Next.Prev = tmpNode
	}
	l.virtualHead.Next = tmpNode
	l.len++
}

func (l *linkedList) AddLast(value interface{}) {
	l.last.Next = NewLinkedNodeByValue(value)
	l.last.Next.Prev = l.last
	l.last = l.last.Next
	l.len++
}

func (l *linkedList) RemoveFirst(obj *interface{}) bool {
	if l.len == 0 || obj == nil {
		return false
	}
	*obj = l.virtualHead.Next.Value

	l.virtualHead.Next = l.virtualHead.Next.Next
	l.virtualHead.Next.Prev = l.virtualHead
	l.len--
	return true
}

func (l *linkedList) RemoveLast(obj *interface{}) bool {
	if l.len == 0 || obj == nil {
		return false
	}
	*obj = l.last.Value

	l.last = l.last.Prev
	l.last.Next = nil
	l.len--
	return true
}

func (l *linkedList) String() string {
	var (
		curNode = l.virtualHead.Next
		builder strings.Builder
	)
	for curNode != nil {
		builder.WriteString(fmt.Sprintf("%v", curNode.Value))
		curNode = curNode.Next
	}
	return builder.String()
}

func (l *linkedList) Len() int {
	return l.len
}

func (l *linkedList) First() *LinkedNode {
	return l.virtualHead.Next
}

func (l *linkedList) Last() *LinkedNode {
	return l.last
}

func NewLinkedList() *linkedList {
	var (
		tmpNode = new(LinkedNode)
	)
	return &linkedList{
		virtualHead: tmpNode,
		last:        tmpNode,
		len:         0,
	}
}
