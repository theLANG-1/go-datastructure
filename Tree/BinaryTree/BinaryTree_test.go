package BinaryTree

import (
	"reflect"
	"testing"

	"github.com/lgbgbl/go-datastructure/Tree/Traversal"
)

func TestNewBinaryTreeByPreAndMid(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		preExpression = "A B D C E"
		midExpression = "D B A C E"
		tree          = NewBinaryTreeByPreAndMid(preExpression, midExpression)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "C", "E"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "A", "C", "E"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByPreAndMid Error")
	}
}
func TestNewBinaryTreeByMidAndPost(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		midExpression  = "D B A C E"
		postExpression = "D B E C A"
		tree           = NewBinaryTreeByMidAndPost(midExpression, postExpression)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "C", "E"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "A", "C", "E"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByMidAndPost Error")
	}
}

func TestNewBinaryTreeByMidAndLevel(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		midExpression   = "D B A C E"
		levelExpression = "A B C D E"
		tree            = NewBinaryTreeByMidAndLevel(midExpression, levelExpression)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "C", "E"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "A", "C", "E"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByMidAndLevel Error")
	}
}

func TestNewBinaryTreeByGListInCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = NewBinaryTreeByGList(expression, Build_GListInCall)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "E", "G", "C", "F"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "G", "E", "A", "C", "F"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByGListInCall Error")
	}
}

func TestNewBinaryTreeByGListInNoCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = NewBinaryTreeByGList(expression, Build_GListInNoCall)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "E", "G", "C", "F"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "G", "E", "A", "C", "F"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByGListInNoCall Error")
	}
}

func TestNewBinaryTreeByPreOrderInCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A B D # # # C # E # #"
		tree       = NewBinaryTreeByPreOrder(expression, "#", Build_PreOrderInCall)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "C", "E"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "A", "C", "E"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByPreOrderInCall Error")
	}
}

func TestNewBinaryTreeByPreOrderInNoCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A B D # # # C # E # #"
		tree       = NewBinaryTreeByPreOrder(expression, "#", Build_PreOrderInNoCall)
	)
	if !reflect.DeepEqual([]interface{}{"A", "B", "D", "C", "E"}, Traversal.PreOrderInCall(tree.Root())) ||
		!reflect.DeepEqual([]interface{}{"D", "B", "A", "C", "E"}, Traversal.MidOrderInCall(tree.Root())) {
		t.Error("NewBinaryTreeByPreOrderInNoCall Error")
	}
}

func BenchmarkNewBinaryTreeByPreAndMid(b *testing.B) {
	var (
		preExpression = "A B D C E"
		midExpression = "D B A C E"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByPreAndMid(preExpression, midExpression)
	}
}

func BenchmarkNewBinaryTreeByMidAndPost(b *testing.B) {
	var (
		midExpression  = "D B A C E"
		postExpression = "D B E C A"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByMidAndPost(midExpression, postExpression)
	}
}

func BenchmarkNewBinaryTreeByMidAndLevel(b *testing.B) {

	var (
		midExpression   = "D B A C E"
		levelExpression = "A B C D E"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByMidAndLevel(midExpression, levelExpression)
	}
}

func BenchmarkNewBinaryTreeByGListInCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByGList(expression, Build_GListInCall)
	}
}

func BenchmarkNewBinaryTreeByGListInNoCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByGList(expression, Build_GListInNoCall)
	}
}

func BenchmarkNewBinaryTreeByPreOrderInCall(b *testing.B) {
	var (
		expression = "A B D # # # C # E # #"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByPreOrder(expression, "#", Build_PreOrderInCall)
	}
}

func BenchmarkNewBinaryTreeByPreOrderInNoCall(b *testing.B) {
	var (
		expression = "A B D # # # C # E # #"
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryTreeByPreOrder(expression, "#", Build_PreOrderInNoCall)
	}
}
