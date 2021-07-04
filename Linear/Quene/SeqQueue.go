package Quene

import (
	"fmt"
	"strings"
)

type seqQuene struct {
	element []interface{}
}

func (s *seqQuene) Enquene(value interface{}) {
	s.element = append(s.element, value)
}

func (s *seqQuene) Dequene(obj *interface{}) bool {
	if !s.First(obj) {
		return false
	}
	s.element = s.element[1:]
	return true
}

func (s *seqQuene) Len() int {
	return len(s.element)
}

func (s *seqQuene) IsEmpty() bool {
	return len(s.element) == 0
}

func (s *seqQuene) String() string {
	var (
		builder strings.Builder
	)
	for _, v := range s.element {
		builder.WriteString(fmt.Sprintf("%v ", v))
	}
	return builder.String()
}

func (s *seqQuene) First(obj *interface{}) bool {
	if s.Len() == 0 {
		return false
	}
	*obj = s.element[0]
	return true
}

func (s *seqQuene) Last(obj *interface{}) bool {
	if s.Len() == 0 {
		return false
	}
	*obj = s.element[s.Len()-1]
	return true
}

func NewSeqQuene() *seqQuene {
	return NewSeqQueneWithCapAndValue(0)
}

func NewSeqQueneWithCapAndValue(capacity int, value ...interface{}) *seqQuene {
	return &seqQuene{
		element: append(make([]interface{}, 0, capacity), value...),
	}
}
