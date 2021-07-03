package Quene

import "testing"

func TestLinkedQueneEnquene(t *testing.T) {
	var (
		q = NewLinkedQuene()
	)
	checkEnquene(t, q)
}

func TestLinkedQueneDequene(t *testing.T) {
	var (
		q = NewLinkedQuene()
	)
	checkDequene(t, q)
}

func BenchmarkLinkedQueneEnquene(b *testing.B) {
	var (
		q = NewLinkedQuene()
	)
	markEnquene(b, q)
}

func TestNewLinkedQueneWithValue(t *testing.T) {
	var (
		elements = []interface{}{1, 2, 3, 4, 5}
		q        = NewLinkedQueneWithValue(elements...)
	)
	checkNewQueneWithValue(t, q, elements)
}
