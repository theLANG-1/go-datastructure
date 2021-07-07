package Tree

type TreeNodeSimplified struct {
	Value  interface{}
	LChild *TreeNodeSimplified
	RChild *TreeNodeSimplified
}

type TreeNode struct {
	Value  interface{}
	LChild *TreeNode
	RChild *TreeNode
	Parent *TreeNode
}

func NewTreeNodeSimplified(value interface{}, lChild, rChild *TreeNodeSimplified) *TreeNodeSimplified {
	return &TreeNodeSimplified{
		Value:  value,
		LChild: lChild,
		RChild: rChild,
	}
}

func NewTreeNodeSimplifiedByValue(value interface{}) *TreeNodeSimplified {
	return NewTreeNodeSimplified(value, nil, nil)
}
