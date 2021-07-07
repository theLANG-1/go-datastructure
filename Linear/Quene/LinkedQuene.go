package Quene

import (
	"fmt"
	"strings"

	"github.com/lgbgbl/go-datastructure/Linear/LinkedList"
)

type linkedQuene struct {
	virtualHead *LinkedList.LinkedNodeSimplified
	last        *LinkedList.LinkedNodeSimplified
	len         int
}

func (l *linkedQuene) Enquene(value interface{}) {
	l.last.Next = LinkedList.NewLinkedNodeSimplifiedByValue(value)
	l.last = l.last.Next
	l.len++
}

func (l *linkedQuene) Dequene(obj *interface{}) bool {
	if !l.First(obj) {
		return false
	}
	l.virtualHead.Next = l.virtualHead.Next.Next
	l.len--
	if l.len == 0 {
		l.last = l.virtualHead
	}
	return true
}

func (l *linkedQuene) First(obj *interface{}) bool {
	if l.Len() == 0 {
		return false
	}
	if obj != nil {
		*obj = l.virtualHead.Next.Value
	}
	return true
}

func (l *linkedQuene) Last(obj *interface{}) bool {
	if l.Len() == 0 {
		return false
	}
	if obj != nil {
		*obj = l.last.Value
	}
	return true
}

func (l *linkedQuene) Len() int {
	return l.len
}

func (l *linkedQuene) IsEmpty() bool {
	return l.len == 0
}

func (l *linkedQuene) String() string {
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

func NewLinkedQuene() *linkedQuene {
	var (
		tmpNode = new(LinkedList.LinkedNodeSimplified)
	)
	return &linkedQuene{
		virtualHead: tmpNode,
		last:        tmpNode,
		len:         0,
	}
}

func NewLinkedQueneWithValue(value ...interface{}) *linkedQuene {
	var (
		newLinkedQuene = NewLinkedQuene()
	)

	for _, v := range value {
		newLinkedQuene.Enquene(v)
	}
	return newLinkedQuene
}
