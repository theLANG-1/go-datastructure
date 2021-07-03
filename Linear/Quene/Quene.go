package Quene

type quene interface {
	Enquene(value interface{})

	Dequene(obj *interface{}) bool

	First(obj *interface{}) bool

	Last(obj *interface{}) bool

	Len() int

	String() string
}
