package Stack

import (
	"fmt"
	"strings"

	"github.com/lgbgbl/go-datastructure/Linear/LinkedList"
)

type linkedStack struct {
	virtualHead *LinkedList.LinkedNodeSimplified
	len         int
}

func (l *linkedStack) Push(value interface{}) {
	l.virtualHead.Next = LinkedList.NewLinkedNodeSimplified(value, l.virtualHead.Next)
	l.len++
}

func (l *linkedStack) Pop(obj *interface{}) bool {
	if !l.Top(obj) {
		return false
	}

	l.virtualHead.Next = l.virtualHead.Next.Next
	l.len--
	return true
}

func (l *linkedStack) Top(obj *interface{}) bool {
	if l.len == 0 {
		return false
	}
	if obj != nil {
		*obj = l.virtualHead.Next.Value
	}
	return true
}

func (l *linkedStack) Len() int {
	return l.len
}

func (l *linkedStack) IsEmpty() bool {
	return l.len == 0
}

func (l *linkedStack) String() string {
	var (
		curNode = l.virtualHead.Next
		builder strings.Builder
	)

	for curNode != nil {
		builder.WriteString(fmt.Sprintf("%v ", curNode.Value))
		curNode = curNode.Next
	}
	return builder.String()
}

func NewLinkedStack() *linkedStack {
	return &linkedStack{
		virtualHead: new(LinkedList.LinkedNodeSimplified),
		len:         0,
	}
}

func NewLinkedStackWithValue(value ...interface{}) *linkedStack {
	var (
		newLinkedStack = NewLinkedStack()
	)

	for _, v := range value {
		newLinkedStack.Push(v)
	}
	return newLinkedStack
}
