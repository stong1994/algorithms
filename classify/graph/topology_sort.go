package graph

/* 拓扑排序：广度优先和贪心算法应用于有向图的专有算法，应用于任务调度、课程安排等
作用：1. 得到一个拓扑序。2. 检测有向图是否有环
https://leetcode-cn.com/problems/course-schedule-ii/solution/tuo-bu-pai-xu-shen-du-you-xian-bian-li-python-dai-/
*/

// 课程安排的合法性
// 你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。
// 例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return false
	}
	// adjs: 被依赖关系——确保由入度为0的节点能够获取到其被依赖的节点
	adjs := make([]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		adjs[i] = make(map[int]struct{})
	}
	// 当前节点的入度：依赖的节点数
	inDegrees := make([]int, numCourses)
	// 初始化入度和依赖关系
	for _, p := range prerequisites {
		adjs[p[1]][p[0]] = struct{}{}
		inDegrees[p[0]]++
	}
	// 模拟栈：存入 入度为0的节点
	var queue []int
	// 先将当前入度为0的节点存入栈
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 从栈中取节点（入度为0的节点），再将节点的临近节点的入度减一，如果临近节点的入度变为0，则将其入栈
	// 统计出栈的节点个数
	cnt := 0
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		cnt++
		for next := range adjs[head] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return cnt == numCourses
}

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
