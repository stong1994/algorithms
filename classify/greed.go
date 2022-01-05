package classify

import (
	"math"
	"sort"
)

/**
贪心算法
贪心算法是最容易理解的一种算法，可能是因为最符合人的思考方式。
运用的关键是了解题意并提出解决办法，然后将流程代码化。
 */


// 分发饼干
// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//对每个孩子 i，都有一个胃口值g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
//并且每块饼干 j，都有一个尺寸 s[j]。如果 s[j]>= g[i]，我们可以将这个饼干 j 分配给孩子 i ，
//这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/assign-cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var result int
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			result++
			i++
			j++
			continue
		}
		j++
	}
	return result
}

// 无重叠区间
// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//注意:
//可以认为区间的终点总是大于它的起点。
//区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-overlapping-intervals
// 找到右端点最小的元素，然后作为第一个元素，再从剩余的元素中找到右端点最小的元素作为第二个元素，以此类推
func eraseOverlapIntervals(intervals [][]int) int {
	var (
		num int
		lo  = math.MinInt64
	)

	sort.Sort(interval(intervals))
	for _, v := range intervals {
		if v[0] >= lo {
			num++
			lo = v[1]
		}
	}
	return len(intervals) - num
}

type interval [][]int

func (i interval) Len() int {
	return len(i)
}

func (i interval) Less(j, k int) bool {
	return i[j][1] < i[k][1]
}

func (i interval) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}

// 用最少数量的箭引爆气球
// 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，
// 因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。
// 一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend，
// 且满足 xstart≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。
// 我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
// 给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons
func findMinArrowShots(points [][]int) int {
	// 根据最右侧元素进行排序，依次干掉剩余元素中右侧元素最小的元素
	sort.Sort(interval(points))
	lo := math.MinInt64
	var result int
	for _, v := range points {
		if v[0] > lo {
			result++
			lo = v[1]
		}
	}
	return result
}

// 根据身高和序号重组队列
// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。
// 每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
// 请你重新构造并返回输入数组people 所表示的队列。返回的队列应该格式化为数组 queue ，
// 其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
// Input:
// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
// Output:
// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
// 思路：高个子不用考虑低个子！所以先按个子从高到低排好，再按照规则将个子低的插入到前边（也就是根据kj来做插入）
func reconstructQueue(people [][]int) [][]int {
	// 先排序：高个子在前，个子相同的按第二个元素个数排序，个数小的排在前
	// 再插入：按照队列中高于等于第j个人的数量为kj来判断j的位置
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	for i, person := range people {
		copy(people[person[1]+1:i+1], people[person[1]:i+1])
		people[person[1]] = person
	}
	return people
}

// 买卖股票最大的收益
// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	// 找到最大值和最小值，要求最小值在最大值的左边
	// 从右往左遍历，保证最小值一定在最大值左边
	min := math.MaxInt64
	var result int
	for _, cur := range prices {
		if cur < min {
			min = cur
		}
		if cur-min > result {
			result = cur - min
		}
	}
	return result
}

// 买卖股票的最大收益 II
// 给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
func maxProfit2(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}

// 种植花朵
// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
// 给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，
// 能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/can-place-flowers
func canPlaceFlowers(flowerbed []int, n int) bool {
	var (
		m int
		l = len(flowerbed)
	)
	for i := 0; i < l; {
		if i == 0 {
			if flowerbed[0] == 1 {
				i += 2
				continue
			}
			if l == 1 {
				m++
				break
			}
			if flowerbed[1] == 0 {
				m++
				i += 2
				continue
			}
			i += 2
			continue
		}
		if i+1 >= l {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				m++
			}
			break
		}

		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				m++
				i += 2
				continue
			}
			i++
			continue
		}
		i += 2
	}
	return m >= n
}

// 观察到需要处理左右两边的边界问题，如果给flowered左右两边各增加一个0元素，即可忽略边界问题！
func canPlaceFlowersOpt(flowerbed []int, n int) bool {
	// 改造flowered
	flowerbed = append([]int{0}, append(flowerbed, 0)...)
	for i := 1; i < len(flowerbed)-1; {
		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				n--
				i += 2
				continue
			}
			i++
			continue
		}
		i += 2
	}
	return n <= 0
}

// 判断是否为子序列
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，
// 而"aec"不是）。
// 进阶：
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/is-subsequence
func isSubsequence(s string, t string) bool {
	lo := 0
	for j := 0; j < len(s); j++ {
		found := false
		for i := lo; i < len(t); i++ {
			if t[i] == s[j] {
				lo = i + 1
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// 修改一个数成为非递减数组
// 给你一个长度为n的整数数组，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。
// 我们是这样定义一个非递减数列的：对于数组中任意的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/non-decreasing-array
func checkPossibility(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	changed := false
	for i := 0; i < len(nums)-1; {
		if nums[i] <= nums[i+1] {
			i++
			continue
		}
		if changed {
			return false
		}
		changed = true
		if i == 0 {
			i++
			continue
		}
		if nums[i+1] >= nums[i-1] {
			nums[i] = nums[i-1]
			continue
		}
		nums[i+1] = nums[i]
		i++
	}
	return true
}

// 子数组最大的和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。
// https://leetcode-cn.com/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		result   = nums[0]
		coResult = nums[0]
	)
	for i := 1; i < len(nums); i++ {
		if coResult <= 0 {
			coResult = nums[i]
		} else {
			coResult += nums[i]
		}
		if coResult > result {
			result = coResult
		}
	}
	return result
}

func maxSubArrayOpt(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		result = nums[0]
	)
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}

// 分隔字符串使同种字符出现在一起
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
// 输入：S = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca", "defegde", "hijhklij"。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/partition-labels
func partitionLabels(s string) []int {
	m := make(map[int32]int) // 字符=》字符首次出现的位置
	list := make([]int, len(s))
	for i, v := range s {
		n, ok := m[v] // 字符在s中的首次出现位置
		if !ok {
			m[v] = i
			list[i] = i
			continue
		}
		list[i] = n
	}
	lastIdx := len(s) - 1
	tmp := lastIdx
	var result []int
	for i := len(s) - 1; i >= 0; i-- { // 从后往前遍历，找到该字符首次出现的位置
		idx := list[i]
		if idx == i && tmp == i {
			result = append(result, lastIdx-i+1)
			lastIdx = i-1
			tmp = i-1
			continue
		}
		if tmp > idx {
			tmp = idx
		}
	}
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - 1 - i
		result[i], result[j] = result[j], result[i]
	}
	return result
}
