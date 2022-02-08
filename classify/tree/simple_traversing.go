package tree

// 二叉树遍历一般有四种遍历方式: 层级遍历、前序遍历、后序遍历、中序遍历
// 层级遍历即广度优先（bfs）,前序遍历、后序遍历、中序遍历都属于深度优先（dfs）
// 前序、后序、中序的差别在于何时处理根节点
// 前序遍历是先处理根节点，再处理左子节点，最后处理右子节点
// 后序遍历是先处理左子节点，再处理右子节点，最后处理根节点
// 中序遍历是先处理左子节点，再处理根节点，最后处理右子节点
// 即前-中-后 是处理根节点的顺序位置

// 二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// 二叉树的后序遍历
// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

// 二叉树的中序遍历
// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}
