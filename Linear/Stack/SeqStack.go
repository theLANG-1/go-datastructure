package Stack

import "fmt"

// 顺序栈
type SeqStack struct {
	element []interface{}
}

func (s *SeqStack) Push(value interface{}) {
	s.element = append(s.element, value)
}

func (s *SeqStack) Pop(obj *interface{}) bool {
	if !s.Top(obj) {
		return false
	}
	s.element = s.element[:len(s.element)-1]
	return true
}

func (s *SeqStack) Len() int {
	return len(s.element)
}

func (s *SeqStack) Top(obj *interface{}) bool {
	if len(s.element) == 0 {
		return false
	}
	if obj != nil {
		*obj = s.element[len(s.element)-1]
	}
	return true
}

func (s *SeqStack) String() string {
	return fmt.Sprintf("%v", s.element)
}

func (s *SeqStack) Cap() int {
	return cap(s.element)
}

func NewSeqStackWithCapAndValue(capacity int, value ...interface{}) *SeqStack {
	return &SeqStack{
		element: append(make([]interface{}, 0, len(value)), value...),
	}
}

func NewSeqStack() *SeqStack {
	return NewSeqStackWithCapAndValue(0)
}
