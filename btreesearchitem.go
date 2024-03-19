package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if ch := BTreeSearchItem(root.Right, elem); ch != nil {
		return ch
	}
	if ch := BTreeSearchItem(root.Left, elem); ch != nil {
		return ch
	}
	return root
}
