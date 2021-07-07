package Traversal

import (
	"reflect"
	"testing"

	"github.com/lgbgbl/go-datastructure/Tree/BinaryTree"
)

func TestPreOrderInNoCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	if tree.Root() == nil {
		t.Error("PreOrderInNoCall Error")
		return
	}
	var (
		preOrderInCallResult   = PreOrderInCall(tree.Root())
		preOrderInNoCallResult = PreOrderInNoCall(tree.Root())
	)
	if !reflect.DeepEqual(preOrderInCallResult, preOrderInNoCallResult) {
		t.Error("PreOrderInNoCall Error")
	}
}

func TestMidOrderInNoCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	if tree.Root() == nil {
		t.Error("MidOrderInNoCall Error")
		return
	}
	var (
		midOrderInCallResult   = MidOrderInCall(tree.Root())
		midOrderInNoCallResult = MidOrderInNoCall(tree.Root())
	)
	if !reflect.DeepEqual(midOrderInCallResult, midOrderInNoCallResult) {
		t.Error("MidOrderInNoCall Error")
		return
	}
}

func TestPostOrderInNoCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	if tree.Root() == nil {
		t.Error("PostOrderInNoCall Error")
		return
	}
	var (
		postOrderInCallResult   = PostOrderInCall(tree.Root())
		postOrderInNoCallResult = PostOrderInNoCall(tree.Root())
	)
	if !reflect.DeepEqual(postOrderInCallResult, postOrderInNoCallResult) {
		t.Error("PostOrderInNoCall Error")
		return
	}
}

func TestLevelOrder(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	if tree.Root() == nil {
		t.Error("LevelOrder Error")
		return
	}

	if !reflect.DeepEqual([]interface{}{"A", "B", "C", "D", "E", "F", "G"}, LevelOrder(tree.Root())) {
		t.Error("LevelOrder Error")
		return
	}
}

func BenchmarkPreOrderInCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreOrderInCall(tree.Root())
	}
}

func BenchmarkPreOrderInNoCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreOrderInNoCall(tree.Root())
	}
}

func BenchmarkMidOrderInCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MidOrderInCall(tree.Root())
	}
}

func BenchmarkMidOrderInNoCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MidOrderInNoCall(tree.Root())
	}
}

func BenchmarkPostOrderInCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostOrderInCall(tree.Root())
	}
}

func BenchmarkPostOrderInNoCall(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostOrderInNoCall(tree.Root())
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	var (
		expression = "A ( B ( D , E ( G , ) ) , C ( , F ) )"
		tree       = BinaryTree.NewBinaryTreeByGList(expression, BinaryTree.Build_GListInNoCall)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LevelOrder(tree.Root())
	}
}
