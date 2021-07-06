package LinkedList

import (
	"math/rand"
	"testing"
	"time"
)

func TestAddFirst(t *testing.T) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		list     = NewLinkedList()
		i        int
	)
	for _, v := range elements {
		list.AddFirst(v)
	}
	if list.Len() != len(elements) {
		t.Error("AddFirst Error")
		return
	}

	first := list.First()
	for i = len(elements) - 1; i >= 0 && first != nil; i-- {
		if elements[i] != first.Value {
			t.Error("AddFirst Error")
			return
		}
		first = first.Next
	}

	if first != nil || i != 0 {
		t.Log("AddFirst Error")
		return
	}
	t.Log("AddFirst Succeed")

}

func BenchmarkAddFirst(b *testing.B) {
	var (
		element = 1
		list    = NewLinkedList()
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddFirst(element)
	}
}

func TestAddLast(t *testing.T) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		list     = NewLinkedList()
		i        int
	)
	for _, v := range elements {
		list.AddLast(v)
	}

	if list.Len() != len(elements) {
		t.Error("AddLast Error")
		return
	}

	first := list.First()
	for i = 0; i < len(elements) && first != nil; i++ {
		if elements[i] != first.Value {
			t.Error("AddLast Error")
			return
		}
		first = first.Next
	}

	if first != nil || i != 0 {
		t.Log("AddLast Error")
		return
	}
	t.Log("AddLast Succeed")

}

func BenchmarkAddLast(b *testing.B) {
	var (
		element = 1
		list    = NewLinkedList()
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddLast(element)
	}
}

func TestRemoveFirst(t *testing.T) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		list     = NewLinkedList()
		target   = new(interface{})
		i        int
		ok       bool
	)

	for _, v := range elements {
		list.AddLast(v)
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(elements)-1) + 1
	for i = 0; i < n; i++ {
		if ok = list.RemoveFirst(target); !ok || *target != elements[i] {
			t.Error("RemoveFirst Fail")
			return
		}
	}
	if list.Len() != len(elements)-n {
		t.Error("RemoveFirst Fail")
	} else if list.First() == nil && n == len(elements) || list.First() != nil && list.First().Value == elements[n] {
		t.Log("RemoveFirst Succeed")
	} else {
		t.Error("RemoveFirst Fail")
	}
}

func TestRemoveLast(t *testing.T) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		list     = NewLinkedList()
		target   = new(interface{})
		i        int
		ok       bool
	)

	for _, v := range elements {
		list.AddLast(v)
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(elements)-1) + 1
	for i = 0; i < n; i++ {
		if ok = list.RemoveLast(target); !ok || *target != elements[len(elements)-i-1] {
			t.Error("RemoveLast Fail")
			return
		}
	}
	if list.Len() != len(elements)-n {
		t.Error("RemoveLast Fail")
	} else if list.Last() == nil && n == 0 || list.Last() != nil && list.Last().Value == elements[len(elements)-n-1] {
		t.Log("RemoveLast Succeed")
	} else {
		t.Error("RemoveLast Fail")
	}
}
