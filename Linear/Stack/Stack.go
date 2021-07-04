package Stack

type stack interface {

	// push an element to stack
	Push(value interface{})

	// pop an element from stack, and obj hold the address of the element
	// return true if the length of stack is not zero, otherwise return false
	Pop(obj *interface{}) bool

	// obj will hold the address of the element if
	// the return value is true
	Top(obj *interface{}) bool

	Len() int

	String() string

	IsEmpty() bool
}
