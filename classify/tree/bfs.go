package tree

// 二叉树的层平均值
// 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10^-5 以内的答案可以被接受
func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	var bfs func(nodes []*TreeNode)
	bfs = func(nodes []*TreeNode) {
		next := make([]*TreeNode, 0, len(nodes)*2)
		val, num := 0, 0
		for _, node := range nodes {
			if node == nil {
				continue
			}
			val += node.Val
			num++
			next = append(next, node.Left, node.Right)
		}
		if num > 0 {
			result = append(result, float64(val)/float64(num))
			bfs(next)
		}
	}
	bfs([]*TreeNode{root})
	return result
}

// 找树左下角的值
// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
// 假设二叉树中至少有一个节点。
func findBottomLeftValue(root *TreeNode) int {
	var result int
	var bfs func(nodes []*TreeNode)
	bfs = func(nodes []*TreeNode) {
		if len(nodes) == 0 {
			return
		}
		result = nodes[0].Val
		next := make([]*TreeNode, 0, len(nodes)*2)
		for _, node := range nodes {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		bfs(next)
	}
	if root == nil {
		return 0
	}
	bfs([]*TreeNode{root})
	return result
}
