package Stack

import "testing"

func TestSeqStackPush(t *testing.T) {
	var (
		s = NewSeqStack()
	)
	checkPush(t, s)
}

func BenchmarkSeqStackPush(b *testing.B) {
	var (
		s = NewSeqStack()
	)
	markStackPush(b, s)
}

func TestSeqStackPop(t *testing.T) {
	var (
		s = NewSeqStack()
	)
	checkPop(t, s)
}

func TestNewSeqStackWithCapAndValue(t *testing.T) {
	var (
		elements = []interface{}{1, 2, 3, 4, 5}
		s        = NewSeqStackWithCapAndValue(0, elements...)
	)
	checkNewStackWithValue(t, s, elements)
}
