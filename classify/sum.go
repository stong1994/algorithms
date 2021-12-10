package classify

import "sort"

// 求和问题往往是N个数相加和为0

// 两数相加: 通过一个数来找到另外一个数，用hash即可
// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, v := range nums {
		hash[v] = i
	}
	for i, v := range nums {
		t := target - v
		if n, ok := hash[t]; ok && n != i {
			return []int{i, n}
		}
	}
	return nil
}

// 三数相加：https://leetcode-cn.com/problems/3sum/
// 思路1：通过一个数找到另外两个数，将所有的数两两组合将和作为key存入map，value为其索引数组。
// 		再遍历一次，通过map来找到符合条件的数据
//		整体时间复杂度=O(N*N+N)
// 思路2：先将所有的数排序，遍历每个数字，然后指定两个指针分别位于首个和末尾的数字，通过比较大小向内移动两个指针
// 		整体时间复杂度=O(N*log(N)+N*N)=O(N^2)
// 综上：思路2的算法更加稳定，思路1如果要考虑去重等操作还需要进一步的运算
func threeSum(nums []int) [][]int {
	var (
		res    [][]int
		length = len(nums)
	)

	sort.Ints(nums)

	for idx := 1; idx < length-1; idx++ {
		start, end := 0, length-1
		// if value of idx and idx-1 equals, then no need to compare all again
		if nums[idx] == nums[idx-1] && idx > 1 {
			start = idx - 1
		}
		for idx > start && idx < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[start] + nums[end] + nums[idx]
			if sum == 0 {
				res = append(res, []int{nums[start], nums[idx], nums[end]})
				start++
				end--
				continue
			}
			if sum < 0 {
				start++
				continue
			}
			end--
		}
	}

	return res
}

// 四数相加
// https://leetcode-cn.com/problems/4sum/
// 排序+双指针：先固定好前两个值，然后通过双指针法找到另外两个值，然后再移动前两个值，完成查找
func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

// https://leetcode-cn.com/problems/4sum-ii/
// 数据已经被“划分”好，通过map两两比较即可
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var totalCount int
	sumMap := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			sumMap[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			totalCount += sumMap[-v3-v4]
		}
	}
	return totalCount
}

// 两数相加: https://leetcode-cn.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		needAdd1    bool
		currentNode *ListNode
		result      *ListNode
	)
	for l1 != nil || l2 != nil {
		val := 0
		if l1 == nil {
			val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		if needAdd1 {
			val++
		}
		if val >= 10 {
			val = val - 10
			needAdd1 = true
		} else {
			needAdd1 = false
		}
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		if currentNode == nil {
			currentNode = node
			result = currentNode
		} else {
			currentNode.Next = node
			currentNode = node
		}
	}
	if needAdd1 {
		currentNode.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return result
}
