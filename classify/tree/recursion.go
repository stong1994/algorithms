package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 树的高度
// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
func maxDepth(root *TreeNode) int {
	// DFS
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 平衡树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/balanced-binary-tree
func isBalanced(root *TreeNode) bool {
	// 是每个节点的左右子树的高度差不超过1，而不是整棵树的高度。因此可以用递归检查每个子树
	// 在获取最长路径的基础上比较左右子树的高度差是否小于1即可
	var getHeight func(root *TreeNode) int // 返回为-1，则表明子树不平衡
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		rh := getHeight(root.Right)
		if rh == -1 {
			return -1
		}
		lh := getHeight(root.Left)
		if lh == -1 {
			return -1
		}
		if rh-lh > 1 || lh-rh > 1 {
			return -1
		}
		return max(rh, lh) + 1
	}
	return getHeight(root) != -1
}

// 两节点的最长路径
// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
// 注意：两结点之间的路径长度是以它们之间边的数目表示。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
func diameterOfBinaryTree(root *TreeNode) int {
	// 一个节点的路径等于左子树的最长路径+右子树的最长路径+1
	// 在递归的过程中，比较当前的路径是否为最长路径
	// 问题即变为寻找最长路径的变种：获取每个节点的高度时，要判断左子树的长度+右子树的长度+1是否为最长的路径
	var result int
	var getHeight func(root *TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := getHeight(root.Left)
		rh := getHeight(root.Right)
		if result < lh+rh {
			result = lh + rh
		}
		return max(lh, rh) + 1
	}

	getHeight(root)
	return result
}

// 翻转树
// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/invert-binary-tree
func invertTree(root *TreeNode) *TreeNode {
	// 问题：对于节点的每个左右子节点都翻转
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 归并两棵树
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。
// 你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-binary-trees
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var left1, left2, right1, right2 *TreeNode
	var val int
	if root1 != nil {
		left1 = root1.Left
		right1 = root1.Right
		val += root1.Val
	}
	if root2 != nil {
		left2 = root2.Left
		right2 = root2.Right
		val += root2.Val
	}
	// 合并每个子树
	return &TreeNode{
		Val:   val,
		Left:  mergeTrees(left1, left2),
		Right: mergeTrees(right1, right2),
	}
}

//  判断路径和是否等于一个数
// 给你二叉树的根节点root 和一个表示目标和的整数targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/path-sum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 统计路径和等于一个数的路径数量
// 给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/path-sum-iii
func pathSum(root *TreeNode, targetSum int) int {
	// 方法1：递归
	// return pathSum1(root, targetSum)
	// 方法2：前缀和：将各个节点的前缀和存入map，在找符合条件的路径时，
	// 只需要将当前节点的前缀和减去targetSum即可得到所需的“前缀”——去掉这个前缀就是我们要的路径
	var result int
	prefixMap := make(map[int]int) // 前缀和=》个数
	prefixMap[0] = 1               // sum-targetSum为0时，表示从该节点到根节点为目标路径
	var recur func(node *TreeNode, subSum int)
	recur = func(node *TreeNode, subSum int) {
		if node == nil {
			return
		}
		sum := node.Val + subSum
		result += prefixMap[sum-targetSum]
		prefixMap[sum]++
		recur(node.Left, sum)
		recur(node.Right, sum)
		prefixMap[sum]-- // 回滚，防止影响兄弟节点
	}
	recur(root, 0)
	return result
}

func pathSum1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 路径必须是向下的，即只能“单向”，对一个节点进行两个方向的遍历
	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + hasSum(root, targetSum)
}

func hasSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	n := 0
	if root.Val == targetSum {
		n++
	}
	return hasSum(root.Left, targetSum-root.Val) + hasSum(root.Right, targetSum-root.Val) + n
}

// 另一个棵树的子树
// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
// 如果存在，返回 true ；否则，返回 false 。
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subtree-of-another-tree
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return isSame(node1.Left, node2.Left) && isSame(node1.Right, node2.Right)
}

// 方法二：可转化为“判断是否是子串”，用KMP实现。根据题意，“子序列”需要保证subRoot的子节点为null时，
// root对应的子节点也要为null，因此，对于为空的子节点要用null来填充
func isSubtree_KMP(root *TreeNode, subRoot *TreeNode) bool {
	var list1, list2 []int
	rootList := tree2list(root, list1)
	subList := tree2list(subRoot, list2)
	return kmp(rootList, subList)
}

func kmp(a, b []int) bool {
	// 构建next数组(由pmt数组优化得)
	next := makeNext(b)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if j == -1 || a[i] == b[j] {
			i++
			j++
			continue
		}
		j = next[j]
	}
	return j == len(b)
}

func makeNext(data []int) []int {
	result := make([]int, len(data))
	result[0] = -1
	cur, pat := 0, -1
	for cur < len(data)-1 {
		if pat == -1 || data[cur] == data[pat] {
			cur++
			pat++
			result[cur] = pat
		} else {
			pat = result[pat]
		}
	}
	return result
}

// 采用中序遍历构建list，左右空节点应用不同的值标识
func tree2list(tree *TreeNode, list []int) []int {
	if tree == nil {
		return list
	}
	list = append(list, tree.Val)
	if tree.Left != nil {
		list = tree2list(tree.Left, list)
	} else {
		list = append(list, math.MinInt32-1)
	}
	if tree.Right != nil {
		list = tree2list(tree.Right, list)
	} else {
		list = append(list, math.MinInt32-2)
	}
	return list
}

// 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//提示：
//树中节点数目在范围 [1, 1000] 内
//-100 <= Node.val <= 100
//进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/symmetric-tree
func isSymmetric(root *TreeNode) bool {
	// 两颗对称树的特性：根节点值相同，左右节点相反。因此可以直接使用递归
	var equal func(a, b *TreeNode) bool
	equal = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		return equal(a.Left, b.Right) && equal(a.Right, b.Left)
	}
	if root == nil {
		return false
	}
	return equal(root.Left, root.Right)
}

// 二叉树的最小深度
// 给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明：叶子节点是指没有子节点的节点。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
func minDepth(root *TreeNode) int {
	// 最小深度-广度优先
	var bp func(nodes []*TreeNode, depth int) int
	bp = func(nodes []*TreeNode, depth int) int {
		next := make([]*TreeNode, 0, len(nodes)*2)
		for _, node := range nodes {
			if node == nil {
				continue
			}
			if node.Left == nil && node.Right == nil {
				return depth
			}
			next = append(next, node.Left, node.Right)
		}
		return bp(next, depth+1)
	}
	if root == nil {
		return 0
	}
	return bp([]*TreeNode{root}, 1)
}

// 左叶子之和
// 计算给定二叉树的所有左叶子之和。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-left-leaves
func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	var getLeft func(node *TreeNode, isLeft bool)
	getLeft = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			result += node.Val
			return
		}
		getLeft(node.Left, true)
		getLeft(node.Right, false)
	}
	getLeft(root, false)
	return result
}

// 最长同值路径
// 给定一个二叉树的root，返回最长的路径的长度 ，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
//两个节点之间的路径长度由它们之间的边数表示。
// [1,null,1,1,1,1,1,1] 的结果为4
//          1
//     nil     	  1
//           1 			1
//		1		1	1
// [1,2,2,2,2,2,2,2]  的结果为3
//			1
//		2		2
// 	 2	 2	  2	  2
// 2
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-univalue-path
func longestUnivaluePath(root *TreeNode) int {
	// 只允许存在一个根节点，而只有根节点允许子节点“分叉”，也就是只能“分叉”一次
	// 每个节点都可能是根节点，因此需要保留其分叉的结果和未分叉的结果，供父节点进行比较、筛选
	var result1 int
	var longPath func(node *TreeNode) [2]int
	longPath = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{}
		}
		maxVal1, maxVal2 := 0, 0
		left := longPath(node.Left)
		right := longPath(node.Right)
		if node.Left != nil && node.Right != nil && node.Left.Val == node.Val && node.Right.Val == node.Val {
			maxVal2 = left[0] + right[0] + 2
		}
		maxVal2 = max(maxVal2, max(left[1], right[1]))
		if node.Left != nil && node.Left.Val == node.Val {
			maxVal1 = left[0] + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			maxVal1 = max(right[0]+1, maxVal1)
		}
		result1 = max(result1, maxVal1)
		return [2]int{maxVal1, maxVal2}
	}
	result := longPath(root)
	return max(result1, result[1])
}

func longestUnivaluePath2(root *TreeNode) int {
	// 需要在递归过程中比较/生成最大值，递归过程中传递的值为当前节点作为单路径节点的最大长度
	var result int
	var dfs func(node *TreeNode, parentVal int) int
	dfs = func(node *TreeNode, parentVal int) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left, node.Val)
		right := dfs(node.Right, node.Val)
		result = max(result, left+right)
		if node.Val == parentVal {
			return max(left, right) + 1
		}
		return 0
	}
	if root == nil {
		return 0
	}
	dfs(root, root.Val)
	return result
}

// 打家劫舍3
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为root。
//除了root之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//给定二叉树的root。返回在不触动警报的情况下，小偷能够盗取的最高金额。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber-iii
func rob(root *TreeNode) int {
	// 每个节点都可以选择偷还是不偷,其状态影响到父节点和子节点
	var dfs func(node *TreeNode) [2]int
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// 选择偷这个节点，那么子节点一定不能偷
		steal := node.Val + left[0] + right[0]
		// 如果不偷这个节点，那么子节点可以偷也可以不偷，选择最大值
		noSteal := max(left[0], left[1]) + max(right[0], right[1])
		return [2]int{noSteal, steal}
	}
	result := dfs(root)
	return max(result[0], result[1])
}

// 二叉树中第二小的节点
// 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。
// 如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
// 更正式地说，即root.val = min(root.left.val, root.right.val) 总成立。
// 给出这样的一个二叉树，你需要输出所有节点中的第二小的值 。
//如果第二小的值不存在的话，输出 -1 。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
func findSecondMinimumValue(root *TreeNode) int {
	// 根据题意，父节点一定是子节点中最小的那个，因此root节点就是最小值，只需要在其子孙节点中找到一个比它大的值即可
	// 但是第二小的值不一定比第三小的值“高",因此不能用广度优先，而需要用深度优先
	if root == nil {
		return -1
	}
	rootVal := root.Val
	result := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val != rootVal && (result == -1 || node.Val < result) {
			result = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}
