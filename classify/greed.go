package classify

import (
	"math"
	"sort"
)

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
	for i, j := 0, 0; i <len(g) && j < len(s); {
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
		lo = math.MinInt64
	)

	sort.Sort(interval(intervals))
	for _, v := range intervals {
		if v[0] >= lo {
			num++
			lo = v[1]
		}
	}
	return len(intervals)- num
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
		if v[0]> lo {
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
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
func reconstructQueue(people [][]int) [][]int {

}