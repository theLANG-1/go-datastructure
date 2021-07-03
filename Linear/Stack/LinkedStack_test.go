package Stack

import "testing"

func TestLinkedStackPush(t *testing.T) {
	var (
		s = NewLinkedStack()
	)
	checkPush(t, s)
}

func BenchmarkLinkedStackPush(b *testing.B) {
	var (
		s = NewLinkedStack()
	)
	markStackPush(b, s)
}

func TestLinkedStackPop(t *testing.T) {
	var (
		s = NewLinkedStack()
	)
	checkPop(t, s)
}

func TestNewLinkedStackWithValue(t *testing.T) {
	var (
		elements = []interface{}{1, 2, 3, 4, 5}
		s        = NewLinkedStackWithValue(elements...)
	)
	checkNewStackWithValue(t, s, elements)
}
