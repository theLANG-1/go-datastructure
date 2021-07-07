package Traversal

import (
	"github.com/lgbgbl/go-datastructure/Linear/Quene"
	"github.com/lgbgbl/go-datastructure/Linear/Stack"
	"github.com/lgbgbl/go-datastructure/Tree"
)

func PreOrderInCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	var (
		preOrder func(rt *Tree.TreeNodeSimplified)
	)
	preOrder = func(rt *Tree.TreeNodeSimplified) {
		if rt != nil {
			result = append(result, rt.Value)
			preOrder(rt.LChild)
			preOrder(rt.RChild)
		}
	}
	preOrder(root)
	return result
}

func MidOrderInCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	var (
		midOrder func(rt *Tree.TreeNodeSimplified)
	)
	midOrder = func(rt *Tree.TreeNodeSimplified) {
		if rt != nil {
			midOrder(rt.LChild)
			result = append(result, rt.Value)
			midOrder(rt.RChild)
		}
	}
	midOrder(root)
	return result
}

func PostOrderInCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	var (
		postOrder func(rt *Tree.TreeNodeSimplified)
	)
	postOrder = func(rt *Tree.TreeNodeSimplified) {
		if rt != nil {
			postOrder(rt.LChild)
			postOrder(rt.RChild)
			result = append(result, rt.Value)
		}
	}
	postOrder(root)
	return result
}

func LevelOrder(root *Tree.TreeNodeSimplified) (result []interface{}) {
	if root == nil {
		return
	}

	var (
		current *Tree.TreeNodeSimplified
		q       = Quene.NewLinkedQueneWithValue(root)
		target  = new(interface{})
	)

	for !q.IsEmpty() {
		q.Dequene(target)
		current = Tree.ToTreeNodeSimPlifiedPointer(target)
		result = append(result, current.Value)
		if current.LChild != nil {
			q.Enquene(current.LChild)
		}
		if current.RChild != nil {
			q.Enquene(current.RChild)
		}
	}
	return
}

func PreOrderInNoCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	if root == nil {
		return
	}
	var (
		s       = Stack.NewLinkedStackWithValue(root)
		target  = new(interface{})
		current *Tree.TreeNodeSimplified
	)
	for !s.IsEmpty() {
		s.Pop(target)
		current = Tree.ToTreeNodeSimPlifiedPointer(target)
		result = append(result, current.Value)
		if current.RChild != nil {
			s.Push(current.RChild)
		}
		if current.LChild != nil {
			s.Push(current.LChild)
		}
	}
	return
}

func MidOrderInNoCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	if root == nil {
		return
	}
	var (
		s       = Stack.NewLinkedStack()
		current = root
		target  = new(interface{})
	)
	for current != nil || !s.IsEmpty() {
		for current != nil {
			s.Push(current)
			current = current.LChild
		}
		if !s.IsEmpty() {
			s.Pop(target)
			current = Tree.ToTreeNodeSimPlifiedPointer(target)
			result = append(result, current.Value)
			current = current.RChild
		}
	}
	return
}

func PostOrderInNoCall(root *Tree.TreeNodeSimplified) (result []interface{}) {
	if root == nil {
		return
	}
	var (
		nodeStack = Stack.NewLinkedStack()
		rootStack = Stack.NewLinkedStack()
		current   = root
		target    = new(interface{})
		last      *Tree.TreeNodeSimplified
	)

	for !nodeStack.IsEmpty() || current != nil {
		for current != nil {
			nodeStack.Push(current)
			if current.RChild != nil {
				rootStack.Push(current)
			}
			current = current.LChild
		}
		nodeStack.Top(target)
		current = Tree.ToTreeNodeSimPlifiedPointer(target)

		rootStack.Top(target)
		tmpRoot := Tree.ToTreeNodeSimPlifiedPointer(target)
		if rootStack.IsEmpty() || current != tmpRoot {
			result = append(result, current.Value)
			last = current
			nodeStack.Pop(nil)
		} else {
			for !nodeStack.IsEmpty() && last == current.RChild {
				result = append(result, current.Value)
				last = current
				nodeStack.Pop(nil)
				rootStack.Pop(nil)
				if nodeStack.IsEmpty() {
					return
				}
				nodeStack.Top(target)
				current = Tree.ToTreeNodeSimPlifiedPointer(target)
			}
		}
		if current != nil {
			current = current.RChild
		}
	}
	return
}
