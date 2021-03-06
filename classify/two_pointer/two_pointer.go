package two_pointer

// 双指针

/*
1. 在一个环中，快慢指针分别每次走两步和一步，如果起始位置不同，能够保证两者一定会相遇？
由于快指针每次都比慢指针多走一步，因此快指针一定会逐渐逼近慢指针。
当快指针落后慢指针两步时，快指针再走一次就会相遇。
当快指针落后慢指针一步时，快指针再走一次就会超过慢指针一步，此时轮到慢指针移动一步，相遇。
（相当于慢指针不动，快指针每次都走一步，因此一定会相遇）
2. 快指针为什么每次走两步而不是三步、四步？
由上分析，当快指针落后于慢指针时，如果每次走两步，那么在追上时一定会相遇，如果是三步、四步则不一定，甚至可能永远不会相遇。
*/

// 寻找重复数
// 给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-the-duplicate-number
// 参考：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/kuai-man-zhi-zhen-de-jie-shi-cong-damien_undoxie-d/
func findDuplicate(nums []int) int {
	// n+1个长度的数组存在一个位于[1,n]的重复数，对于nums中的元素v，可视为其nums中的索引
	// 对于无重复数的有序数组[1,2,3,4,5]，nums[i] == nums[nums[i-1]]
	// 对于无重复数的无序数组[3,1,4,5,2]，我们可以通过next=next.next的方式来遍历nums，遍历顺序为3,4,5,2,1。如果不终止的话，还能进入第二次循环。
	// 对于有重复数的无序数组[3,1,4,5,2,3]，当我们通过next=next.next的方式遍历到最后一个4时，next=4，因此整个遍历顺序为[3,1,4,5,2][3,4,5,2][3,4,5,2]
	// 发现[3,4,5,2]进入循环，根据上边的经验，我们也能知道循环数组中的第一个元素就是我们需要的重复数
	// 对于有重复数的无序数组，可以看做是一段线+一段环，线与环的交点就是重复元素。
	// 设置快慢两个指针，快指针会先于进入环，然后慢指针进入，由于快指针每次都相遇慢指针多移动一步，因此他们会相遇（相遇时不一定是重复数）。
	// 假设非环部分长度为m，环的长度为k，当两个指针相遇时，慢指针走了n,快指针走了2n,那么多出来的n就是环的长度k的整数倍。
	// 此时慢指针在环中的相对位置为n-m，此时需要再走m就能到达入口，也就是目标循环数
	// 如何求m？设一个新的指针，从头开始移动，每次移动一步，移动到环的入口时正好碰到慢指针。
	slow, fast := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	helper := 0
	for {
		if nums[helper] == nums[slow] {
			return nums[slow]
		}
		helper = nums[helper]
		slow = nums[slow]
	}
}
