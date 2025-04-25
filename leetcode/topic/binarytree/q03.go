package binarytree

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			res = append(res, root.Val)
		}
	}
	dfs(root)
	return res
}
