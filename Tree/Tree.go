package Tree

func ToTreeNodeSimPlifiedPointer(obj *interface{}) *TreeNodeSimplified {
	if obj == nil {
		return nil
	}
	return (*obj).(*TreeNodeSimplified)
}
