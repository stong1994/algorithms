package link_list

// 链表

// 找出两个链表的交点
// 给你两个单链表的头节点headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。
//自定义评测：
//评测系统 的输入如下（你设计的程序 不适用 此输入）：
//intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
//listA - 第一个链表
//listB - 第二个链表
//skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
//skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
//评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。
//如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 题中有两个点：
	// 1. 如果找到一个相同的nodeC，那么对于两个链表nodeC之后的node也一定相同
	// 2. 比较node不能只比较val，要比较指针地址
	/**
	方法1：构建双层循环遍历两个链表，找到相同的node即可，时间复杂度是O(m*n)
	方法2：将两个链表分成3部分：专属于headA的A链表、专属于headB的B链表、同一个链表C。
		在headA后边拼接B，在headB后边拼接A，此时两个拼接后的链表长度相同，那么最后一个node的Next就是链表C的头结点。
		即遍历长度为len(headA)+len(headB)的链表，分别从HeadA和headB开始，如果能找到相同的node，即为目标node，否则表示没有相同的node
	*/
	firstA, firstB := headA, headB
	for {
		if headA == headB {
			return headA
		}
		// 如果两个链表长度相同，此时仍没有找到相同的node，则表示没有相同的node。如果不同，那么此时表示遍历完了链表A和链表B组成的链表，仍没有找到
		if headA.Next == nil && headB.Next == nil {
			return nil
		}
		if headA.Next != nil {
			headA = headA.Next
		} else {
			headA = firstB
		}
		if headB.Next != nil {
			headB = headB.Next
		} else {
			headB = firstA
		}
	}
}

// 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next) // 反转剩余部分的node，
	// 此时head变为head.Next.Next，而head的Next变为nil
	head.Next.Next = head
	head.Next = nil
	return node
}

// 归并两个有序的链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 提示：
//	两个链表的节点数目范围是 [0, 50]
//	-100 <= Node.val <= 100
//	l1 和 l2 均按 非递减顺序 排列
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}

// 从有序链表中删除重复节点
// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
// 返回同样按升序排列的结果链表。
// 提示：
//	链表中节点数目在范围 [0, 300] 内
//	-100 <= Node.val <= 100
//	题目数据保证链表已经按升序排列
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

// 删除链表的倒数第 n 个节点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 方法1：数组/栈：将这些节点依次入栈，然后再出栈，
	// 出栈时计数，即能找到倒数第n个节点，此时栈中剩下的就是倒数n-1个节点，将其Next指针指向刚出栈的节点的下一个节点即可
	//if n <= 0 {
	//	return head
	//}
	//var list []*ListNode
	//node := head
	//for node != nil {
	//	list = append(list, node)
	//	node = node.Next
	//}
	//if len(list) < n {
	//	return head
	//}
	//if len(list) == n {
	//	return head.Next
	//}
	//
	//target := list[len(list)-n-1]
	//target.Next = target.Next.Next
	//return head

	// 方法2：快慢双指针，快指针比慢指针领先n个节点，如实在快节点到达末尾时，慢节点位于倒数第n个节点。
	// 此时的慢节点就是我们要删掉的节点，因此可令n为n+1，此时的慢节点就是要删掉的前一个节点。
	dummy := new(ListNode)
	dummy.Next = head
	node1 := dummy
	node2 := dummy
	for i := 0; i < n+1; i++ {
		if node1 == nil { // 说明链表长度小于n
			return head
		}
		node1 = node1.Next
	}
	for node1 != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	node2.Next = node2.Next.Next
	return dummy.Next
}

// 交换链表中的相邻结点
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 提示：
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	remain := swapPairs(head.Next.Next)
	second := head.Next
	second.Next = head
	head.Next = remain
	return second
}

// 链表求和
// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
// 提示：
//	链表的长度范围为 [1, 100]
//	0 <= node.val <= 9
//	输入数据保证链表代表的数字无前导 0
//	进阶：如果输入链表不能翻转该如何解决？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers-ii
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		list1, list2 []*ListNode
	)
	for l1 != nil {
		list1 = append(list1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		list2 = append(list2, l2)
		l2 = l2.Next
	}
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
	}
	// 对list2填充前置0
	zeroList := make([]*ListNode, len(list1)-len(list2))
	for i := range zeroList {
		zeroList[i] = new(ListNode)
	}
	list2 = append(zeroList, list2...)
	needAdd := 0
	for i := len(list1) - 1; i >= 0; i-- {
		val := list1[i].Val + list2[i].Val + needAdd
		if val >= 10 {
			list1[i].Val = val - 10
			needAdd = 1
		} else {
			list1[i].Val = val
			needAdd = 0
		}
	}
	if needAdd == 1 {
		node := new(ListNode)
		node.Val = 1
		node.Next = list1[0]
		return node
	}
	return list1[0]
}

// 回文链表
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
// 提示：
//	链表中节点数目在范围[1, 105] 内
//	0 <= Node.val <= 9
// 进阶：你能否用O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-linked-list
func isPalindrome(head *ListNode) bool {
	// 1. 通过快慢双指针来找到链表后半部分：快指针走两步，慢指针走一步。
	// 从第一个节点开始走，如果总数为奇数，那么快指针走完时，慢指针走到中间节点的下一个；如果总数为偶数，那么快指针走完时，慢指针走到后半部分的第一个节点
	fastNode, slowNode := head, head
	for fastNode != nil {
		fastNode = fastNode.Next
		if fastNode != nil {
			fastNode = fastNode.Next
		}
		slowNode = slowNode.Next
	}
	// 2. 反转后半部分
	slowNode = reverseList(slowNode)
	// 3. 对前后两半部分比较
	for slowNode != nil {
		if slowNode.Val != head.Val {
			return false
		}
		slowNode = slowNode.Next
		head = head.Next
	}
	return true
}

// 分隔链表
// 给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
// 每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
// 这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
// 返回一个由上述 k 部分组成的数组。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/split-linked-list-in-parts
func splitListToParts(head *ListNode, k int) []*ListNode {
	// 先将节点存入数组，然后把数组分成k份，假设数组长度为L
	// 如果L==k，则每份正好只有1个
	// 如果L<k,则每份最多有1个，先到先得，后边的补nil
	// 如果L>k，则进行整除，假设余数为n，则前n份还要再加1份
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	L := len(list)
	n1 := L / k
	n2 := L - n1*k
	result := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		var (
			firstIdx, lastIdx int
		)
		if i < n2 {
			firstIdx = i * (n1 + 1)
			lastIdx = firstIdx + n1
		} else {
			firstIdx = i*n1 + n2
			lastIdx = firstIdx + n1 - 1
		}
		if firstIdx >= L {
			result[i] = nil
		} else {
			if lastIdx < L {
				list[lastIdx].Next = nil
			}
			result[i] = list[firstIdx]
		}
	}
	return result
}

// 链表元素按奇偶聚集
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
// 说明:
// 应当保持奇数节点和偶数节点的相对顺序。
// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/odd-even-linked-list
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	evenHead := head.Next
	even := evenHead
	odd := head
	for even != nil && even.Next != nil {
		odd.Next = even.Next // 奇数的下一个是偶数的下一个
		odd = odd.Next
		even.Next = odd.Next // 偶数的下一个是奇数的下一个
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
