package tree

import "math"

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

// 从有序数组中构造二叉查找树
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
func sortedArrayToBST(nums []int) *TreeNode {
	// 选择中间节点作为根节点，左边的节点作为左子树，右边的节点作为右子树，并递归
	L := len(nums)
	if L == 0 {
		return nil
	}
	if L == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	return &TreeNode{
		Val:   nums[L/2],
		Left:  sortedArrayToBST(nums[:L/2]),
		Right: sortedArrayToBST(nums[L/2+1:]),
	}
}

// 根据有序链表构造平衡的二叉查找树
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
func sortedListToBST(head *ListNode) *TreeNode {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST2(head *ListNode) *TreeNode {
	// 找到中间节点，作为根节点，剩余左部分做左子树，右部分做右子树
	// 中间节点可使用快慢双指针
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	fast, slow := head, head
	var lastSlow *ListNode
	for fast != nil && fast.Next != nil {
		lastSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	lastSlow.Next = nil
	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST2(head),
		Right: sortedListToBST2(slow.Next),
	}
}

// *****
func sortedListToBST3(head *ListNode) *TreeNode {
	// 链表的值是递增，正好符合中序遍历，因此可以利用中序遍历来构建二叉查找树
	// 对于每棵树，先构建其左子节点，再构建其根节点，然后是右子节点，这样符合链表的遍历顺序，但是需要预留根节点的值，因此使用中序遍历
	L := 0
	h := head
	for head != nil {
		head = head.Next
		L++
	}
	var build func(start, end int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := new(TreeNode)
		root.Left = build(start, mid-1)
		root.Val = h.Val
		h = h.Next
		root.Right = build(mid+1, end)
		return root
	}
	return build(0, L-1)
}

// 在二叉查找树中寻找两个节点，使它们的和为一个给定值
// 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
// 提示:
//二叉树的节点个数的范围是[1, 10^4].
//-10^4<= Node.val <= 10^4
//root为二叉搜索树
//-10^5<= k <= 10^5
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
func findTarget(root *TreeNode, k int) bool {
	// 如果是在一个有序数组中找到和为k的两个元素要如何做？利用分别指向最大值和最小值的双指针，比较和的大小并移动指针即可。
	// 二叉树中只有左子树和右子树，不能同时从两头开始遍历，因此需要同时进行两次遍历，一个先遍历左子树，另一个先遍历右子树，即一个为前序遍历，另一个为逆序的中序遍历
	// 用数量来记录两者最多能够移动的步数
	//var getNum func(node *TreeNode) int
	//getNum = func(node *TreeNode) int {
	//	if node == nil {
	//		return 0
	//	}
	//	return 1 + getNum(node.Left) +getNum(node.Right)
	//}

	var (
		left     int
		right    int
		leftIdx  int
		rightIdx int
	)

	var list []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		list = append(list, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	if len(list) < 2 {
		return false
	}
	leftIdx = 0
	rightIdx = len(list) - 1

	for leftIdx < rightIdx {
		left = list[leftIdx]
		right = list[rightIdx]
		if left+right < k {
			leftIdx++
		} else if left+right > k {
			rightIdx--
		} else {
			return true
		}
	}
	return false
}

// 二叉搜索树的最小绝对差
// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
// 差值是一个正数，其数值等于两值之差的绝对值。
// 提示：
//	树中节点的数目范围是 [2, 10^4]
//	0 <= Node.val <= 10^5
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
func getMinimumDifference(root *TreeNode) int {
	// 二叉查找树可以看做是递增的数组，只要比较相邻的元素大小即可。
	var (
		last   int
		minVal = math.MaxInt32
	)
	initLast := false
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if !initLast {
			last = node.Val
			initLast = true
		} else {
			minVal = min(node.Val-last, minVal)
			last = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 寻找二叉查找树中出现次数最多的值
// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//如果树中有不止一个众数，可以按 任意顺序 返回。
//假定 BST 满足如下定义：
//结点左子树中所含节点的值 小于等于 当前节点的值
//结点右子树中所含节点的值 大于等于 当前节点的值
//左子树和右子树都是二叉搜索树
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree
func findMode(root *TreeNode) []int {
	// 在一个有序数组中，从左到右依次遍历
	curNum := 0
	maxNum := 0
	last := math.MaxInt32
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val != last {
			if curNum > maxNum {
				result = []int{last}
				maxNum = curNum
			} else if curNum == maxNum && curNum != 0 {
				result = append(result, last)
			}
			curNum = 1
			last = node.Val
		} else {
			curNum++
		}
		dfs(node.Right)
	}
	dfs(root)
	if curNum == maxNum {
		result = append(result, last)
	}
	if curNum > maxNum {
		result = []int{last}
	}
	return result
}
