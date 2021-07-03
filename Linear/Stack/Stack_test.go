package Stack

import (
	"math/rand"
	"testing"
	"time"
)

func checkPush(t *testing.T, s stack) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		target   = new(interface{})
		ok       bool
	)
	for _, v := range elements {
		s.Push(v)
	}

	if s.Len() != len(elements) {
		t.Error("Push Error")
		return
	}
	if ok = s.Top(target); !ok || *target != elements[len(elements)-1] {
		t.Error("Push Error")
		return
	}
	t.Log("Push Succeed")
}

func checkPop(t *testing.T, s stack) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		target   = new(interface{})
		ok       bool
	)
	for _, v := range elements {
		s.Push(v)
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(elements)-1) + 1
	for i := 0; i < n; i++ {
		if ok = s.Pop(target); !ok || *target != elements[len(elements)-1-i] {
			t.Error("Pop Error")
			return
		}
	}

	if s.Len() != len(elements)-n {
		t.Error("Pop Error")
		return
	}
	t.Log("Pop Succeed")
}

func checkNewStackWithValue(t *testing.T, s stack, elements []interface{}) {
	var (
		target = new(interface{})
		ok     bool
	)
	if s.Len() != len(elements) {
		t.Error("NewStackWithValue Error")
		return
	}

	for i := 0; i < len(elements); i++ {
		if ok = s.Pop(target); !ok || *target != elements[len(elements)-i-1] {
			t.Error("NewStackWithValue Error")
			return
		}
	}
	t.Log("NewStackWithValue Succeed")
}

func markStackPush(b *testing.B, s stack) {
	var (
		element = 1
	)
	for i := 0; i < b.N; i++ {
		s.Push(element)
	}
}
