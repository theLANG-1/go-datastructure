package Quene

import "testing"

func TestSeqQueneEnquene(t *testing.T) {
	var (
		q = NewSeqQuene()
	)
	checkEnquene(t, q)
}

func TestSeqQueneDequene(t *testing.T) {
	var (
		q = NewSeqQuene()
	)
	checkDequene(t, q)
}

func BenchmarkSeqQueneEnquene(b *testing.B) {
	var (
		q = NewSeqQuene()
	)
	markEnquene(b, q)
}

func TestNewSeqQueneWithCapAndValue(t *testing.T) {
	var (
		elements = []interface{}{1, 2, 3, 4, 5}
		q        = NewSeqQueneWithCapAndValue(0, elements...)
	)
	checkNewQueneWithValue(t, q, elements)
}
