package BinaryTree

import (
	"strings"

	"github.com/lgbgbl/go-datastructure/Linear/Stack"
	"github.com/lgbgbl/go-datastructure/Tree"
)

const (
	Build_PreOrderInCall = iota
	Build_PreOrderInNoCall
	Build_GListInCall
	Build_GListInNoCall
)

type binaryTree struct {
	root *Tree.TreeNodeSimplified
}

func (b *binaryTree) Root() *Tree.TreeNodeSimplified {
	return b.root
}

func NewBinaryTreeByRoot(root *Tree.TreeNodeSimplified) *binaryTree {
	return &binaryTree{
		root: root,
	}
}

func NewBinaryTree() *binaryTree {
	return NewBinaryTreeByRoot(nil)
}

func NewBinaryTreeByPreOrder(pExpression, endSymbol string, flag int) *binaryTree {
	var (
		root   *Tree.TreeNodeSimplified
		tokens = strings.Split(pExpression, " ")
	)
	switch flag {
	case Build_PreOrderInCall:
		index := 0
		root = buildTreeNodeByPreOrderInCall(tokens, &index, endSymbol)
	case Build_PreOrderInNoCall:
		root = buildTreeNodeByPreOrderInNoCall(tokens, endSymbol)
	}
	return NewBinaryTreeByRoot(root)
}

func NewBinaryTreeByGList(expression string, flag int) *binaryTree {
	var (
		root   *Tree.TreeNodeSimplified
		tokens = strings.Split(expression, " ")
	)

	switch flag {

	case Build_GListInCall:
		index := 0
		root = buildTreeNodeByGListInCall(tokens, &index, nil)
	case Build_GListInNoCall:
		root = buildTreeNodeByGListInNoCall(tokens)
	}
	return NewBinaryTreeByRoot(root)
}

func NewBinaryTreeByPreAndMid(pExpression, mExpression string) *binaryTree {
	var (
		pTokens = strings.Split(pExpression, " ")
		mTokens = strings.Split(mExpression, " ")
		root    = buildTreeNodeByPreAndMid(pTokens, mTokens)
	)
	return NewBinaryTreeByRoot(root)
}

func NewBinaryTreeByMidAndPost(mExpression, pExpression string) *binaryTree {
	var (
		mTokens = strings.Split(mExpression, " ")
		pTokens = strings.Split(pExpression, " ")
		root    = buildTreeNodeByMidAndPost(mTokens, pTokens)
	)
	return NewBinaryTreeByRoot(root)
}

func NewBinaryTreeByMidAndLevel(mExpression, lExpression string) *binaryTree {
	var (
		mTokens = strings.Split(mExpression, " ")
		pTokens = strings.Split(lExpression, " ")
		root    = buildTreeNodeByMidAndLevel(mTokens, pTokens)
	)
	return NewBinaryTreeByRoot(root)
}

func buildTreeNodeByPreOrderInCall(tokens []string, indexPointer *int, endSymbol string) *Tree.TreeNodeSimplified {
	var (
		root      *Tree.TreeNodeSimplified
		rootValue = tokens[*indexPointer]
	)

	if rootValue == endSymbol {
		*indexPointer++
	} else {
		root = Tree.NewTreeNodeSimplifiedByValue(rootValue)
		*indexPointer++
		root.LChild = buildTreeNodeByPreOrderInCall(tokens, indexPointer, endSymbol)
		root.RChild = buildTreeNodeByPreOrderInCall(tokens, indexPointer, endSymbol)
	}
	return root
}

func buildTreeNodeByPreOrderInNoCall(tokens []string, endSymbol string) *Tree.TreeNodeSimplified {
	if len(tokens) == 0 {
		return nil
	}
	var (
		lastNode *Tree.TreeNodeSimplified
		root     = Tree.NewTreeNodeSimplifiedByValue(tokens[0])
		s        = Stack.NewLinkedStackWithValue(root)
		target   = new(interface{})
	)
	for i := 1; i < len(tokens); i++ {
		if tokens[i] != endSymbol {
			lChild := Tree.NewTreeNodeSimplifiedByValue(tokens[i])
			s.Top(target)
			lastNode = Tree.ToTreeNodeSimPlifiedPointer(target)
			lastNode.LChild = lChild
			s.Push(lChild)
		} else {
			for i < len(tokens) && tokens[i] == endSymbol {
				s.Pop(target)
				lastNode = Tree.ToTreeNodeSimPlifiedPointer(target)
				i++
			}
			if i >= len(tokens) {
				break
			}

			rChild := Tree.NewTreeNodeSimplifiedByValue(tokens[i])
			lastNode.RChild = rChild
			s.Push(rChild)
		}
	}
	return root
}

func buildTreeNodeByGListInCall(tokens []string, indexPointer *int, parent *Tree.TreeNodeSimplified) (current *Tree.TreeNodeSimplified) {

	if *indexPointer < len(tokens) && tokens[*indexPointer] != "," && tokens[*indexPointer] != ")" && tokens[*indexPointer] != "(" {
		current = Tree.NewTreeNodeSimplifiedByValue(tokens[*indexPointer])
		*indexPointer++
	}
	if *indexPointer < len(tokens) && tokens[*indexPointer] == "(" {
		*indexPointer++
		current.LChild = buildTreeNodeByGListInCall(tokens, indexPointer, current)
	}
	if *indexPointer < len(tokens) && tokens[*indexPointer] == "," {
		*indexPointer++
		parent.RChild = buildTreeNodeByGListInCall(tokens, indexPointer, parent)
	} else if *indexPointer < len(tokens) && tokens[*indexPointer] == ")" {
		*indexPointer++
	}
	return
}

func buildTreeNodeByGListInNoCall(tokens []string) *Tree.TreeNodeSimplified {
	if len(tokens) == 0 {
		return nil
	}

	var (
		root   = Tree.NewTreeNodeSimplifiedByValue(tokens[0])
		in     = root
		out    *Tree.TreeNodeSimplified
		flag   uint8
		s      = Stack.NewLinkedStack()
		target = new(interface{})
	)

	for i := 1; i < len(tokens); i++ {
		switch tokens[i] {
		case "(":
			s.Push(in)
			flag = 1
		case ")":
			s.Pop(nil)
		case ",":
			flag = 2
		default:
			in = Tree.NewTreeNodeSimplifiedByValue(tokens[i])
			s.Top(target)
			out = Tree.ToTreeNodeSimPlifiedPointer(target)
			if flag == 1 {
				out.LChild = in
			} else {
				out.RChild = in
			}
		}
	}

	return root
}

func buildTreeNodeByPreAndMid(pTokens, mTokens []string) *Tree.TreeNodeSimplified {
	if len(pTokens) == 0 {
		return nil
	}

	var (
		root = Tree.NewTreeNodeSimplifiedByValue(pTokens[0])
		pos  int
	)

	for pos = 0; pos < len(pTokens); pos++ {
		if mTokens[pos] == pTokens[0] {
			break
		}
	}
	root.LChild = buildTreeNodeByPreAndMid(pTokens[1:pos+1], mTokens[:pos])
	root.RChild = buildTreeNodeByPreAndMid(pTokens[pos+1:], mTokens[pos+1:])
	return root
}

func buildTreeNodeByMidAndPost(mTokens, pTokens []string) *Tree.TreeNodeSimplified {
	if len(mTokens) == 0 {
		return nil
	}

	var (
		root = Tree.NewTreeNodeSimplifiedByValue(pTokens[len(pTokens)-1])
		pos  int
	)

	for pos = 0; pos < len(mTokens); pos++ {
		if mTokens[pos] == pTokens[len(pTokens)-1] {
			break
		}
	}
	root.LChild = buildTreeNodeByMidAndPost(mTokens[:pos], pTokens[:pos])
	root.RChild = buildTreeNodeByMidAndPost(mTokens[pos+1:], pTokens[pos:len(pTokens)-1])
	return root
}

func buildTreeNodeByMidAndLevel(mTokens, lTokens []string) *Tree.TreeNodeSimplified {
	if len(lTokens) == 0 {
		return nil
	}

	var (
		left, right []string
		pos         int
		root        = Tree.NewTreeNodeSimplifiedByValue(lTokens[0])
	)

	for k, v := range mTokens {
		if v == lTokens[0] {
			pos = k
			break
		}
	}

	for i := 1; i < len(lTokens); i++ {
		if existIn(lTokens[i], mTokens[:pos]) {
			left = append(left, lTokens[i])
		} else {
			right = append(right, lTokens[i])
		}
	}

	root.LChild = buildTreeNodeByMidAndLevel(mTokens[:pos], left)
	root.RChild = buildTreeNodeByMidAndLevel(mTokens[pos+1:], right)
	return root
}

func existIn(element string, elements []string) bool {
	for _, v := range elements {
		if element == v {
			return true
		}
	}
	return false
}
