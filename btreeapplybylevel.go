package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	tmp := []*TreeNode{root}

	for len(tmp) != 0 {
		c := tmp[0]
		tmp = tmp[1:]

		_, err := f(c.Data)
		if err != nil {
			return
		}

		if c.Left != nil {
			tmp = append(tmp, c.Left)
		}
		if c.Right != nil {
			tmp = append(tmp, c.Right)
		}
	}
}
