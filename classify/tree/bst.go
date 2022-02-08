package tree

// 二叉查找树
// 二叉查找树的特点：根节点大于等于左子树所有节点，小于等于右子树所有节点——因此中序遍历是非递减的。
// 修剪二叉搜索树
// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
// 修剪树 不应该改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在唯一的答案。
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
// 提示：
//	树中节点数在范围 [1, 10^4] 内
//	0 <= Node.val <= 10^4
//	树中每个节点的值都是 唯一 的
//	题目数据保证输入是一棵有效的二叉搜索树
//	0 <= low <= high <= 10^4
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/trim-a-binary-search-tree
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// 遍历树，移除边界外的节点
	// 当移除节点时，如果是小于边界，则其左子树的值都小于边界，因此只保留右子树；相反，只保留左子树
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < low {
			return dfs(node.Right)
		}
		if node.Val > high {
			return dfs(node.Left)
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		return node
	}
	return dfs(root)
}

// 	二叉搜索树中第K小的元素
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历，找到第k个遍历的值
	val, num := 0, 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		num++
		if num == k {
			val = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return val
}

// 把二叉搜索树转换为累加树
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node的新值等于原树中大于或等于node.val的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree
func convertBST(root *TreeNode) *TreeNode {
	// 先遍历右子节点，再处理当前值，最后处理左子节点。每个节点的值都等于右子节点的值+当前最大值
	curMax := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += curMax
		curMax = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}

// 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
//满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点的值大于p的值和q的值，说明祖先节点在当前节点的右子树上
	// 如果当前节点的值小于p和值和q的值，说明祖先节点在当前节点的左子树上
	// 否则，说明当前节点就是p和q的祖先节点
	var (
		result *TreeNode
		pv     = p.Val
		qv     = q.Val
	)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		v := node.Val
		if v > pv && v > qv {
			dfs(node.Left)
			return
		}
		if v < pv && v < qv {
			dfs(node.Right)
			return
		}
		result = node
	}
	dfs(root)
	return result
}

// 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 根据题意 p和q都存在树中
	// 对于一个节点来说，有几种可能：
	// p和q分别在左子树节点和右子树节点，则当前节点就是祖先节点
	// p和q同时在一颗子树，那么p或者q就是祖先节点
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
