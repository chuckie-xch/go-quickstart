package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		arr = append(arr, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return arr
}

func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	st := []*TreeNode{}
	st = append(st, root)
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
	return res
}
