package graph

/* 拓扑排序：广度优先和贪心算法应用于有向图的专有算法，应用于任务调度、课程安排等
作用：1. 得到一个拓扑序。2. 检测有向图是否有环
*/

// 课程安排的顺序
// 现在你总共有 numCourses 门课需要选，记为0到numCourses - 1。给你一个数组prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修bi 。
// 例如，想要学习课程 0 ，你需要先完成课程1 ，我们用一个匹配来表示：[0,1] 。
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule-ii
func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return nil
	}
	adjs := make([]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		adjs[i] = make(map[int]struct{})
	}
	inDegree := make([]int, numCourses) // 入度：依赖的节点数
	for _, p := range prerequisites {
		adjs[p[1]][p[0]] = struct{}{}
		inDegree[p[0]]++
	}
	queue := make([]int, 0, numCourses) // 模拟堆栈
	// 先将入度为0的节点push到堆中
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	res := make([]int, numCourses)
	count := 0 // 已安排好的课程数，正好可以作为索引
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		res[count] = head
		count++
		successors := adjs[head]
		for nextCourse := range successors {
			inDegree[nextCourse]--
			// 检测节点的入度是否为0，如果是，则加入队列
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}
	// 如果结果集中的数量不等于节点的数量，就不能完成课程任务
	if count != numCourses {
		return nil
	}
	return res
}
