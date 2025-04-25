package binarytree

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			res = append(res, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
