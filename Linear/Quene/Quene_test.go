package Quene

import (
	"math/rand"
	"testing"
	"time"
)

func checkEnquene(t *testing.T, q quene) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		target   = new(interface{})
		ok       bool
	)
	for _, v := range elements {
		q.Enquene(v)
	}

	if q.Len() != len(elements) {
		t.Error("Enquene Error")
		return
	}

	ok = q.First(target)
	if !ok || *target != elements[0] {
		t.Error("Enquene Error")
		return
	}

	ok = q.Last(target)
	if !ok || *target != elements[len(elements)-1] {
		t.Error("Enquene Error")
		return
	}
}

func markEnquene(b *testing.B, q quene) {
	var (
		element = 1
	)
	for i := 0; i < b.N; i++ {
		q.Enquene(element)
	}
}

func checkDequene(t *testing.T, q quene) {
	var (
		elements = []int{1, 2, 3, 4, 5}
		target   = new(interface{})
		ok       bool
	)
	for _, v := range elements {
		q.Enquene(v)
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(elements)-1) + 1
	for i := 0; i < n; i++ {
		if ok = q.Dequene(target); !ok || *target != elements[i] {
			t.Error("Dequene Error")
			return
		}
	}

	if q.Len() != len(elements)-n {
		t.Error("Dequene Error")
		return
	}
	t.Log("Dequene Succeed")
}

func checkNewQueneWithValue(t *testing.T, q quene, elements []interface{}) {
	var (
		target = new(interface{})
		ok     bool
	)

	if q.Len() != len(elements) {
		t.Error("NewQueneWithValue Error")
		return
	}

	for i := 0; i < len(elements); i++ {
		if ok = q.Dequene(target); !ok || *target != elements[i] {
			t.Error("NewQueneWithValue Error")
			return
		}
	}
	t.Log("NewQueneWithValue Succeed")
}
