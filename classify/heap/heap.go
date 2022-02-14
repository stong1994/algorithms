package heap

// 最大堆
type MaxHeap struct {
	pq []int // 优先队列 首个元素不存储值，对于父节点索引i，其子节点的索引分别为 2*i、2*i+1
	n  int   // 元素个数，则 n == len(pq)-1
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		pq: []int{0},
		n:  0,
	}
}

func (mh *MaxHeap) isEmpty() bool {
	return mh.n == 0
}

// 删除最大元素：删除堆顶元素，并将末尾元素放到堆顶，对堆顶元素进行下沉
func (mh *MaxHeap) delMax() int {
	if mh.isEmpty() {
		return 0
	}
	max := mh.pq[1]
	mh.pq[1] = mh.pq[mh.n]
	mh.n--
	mh.pq = mh.pq[:mh.n+1]
	mh.sink(1)
	return max
}

// 插入元素，将元素放到末尾，然后进行上浮
func (mh *MaxHeap) insert(k int) {
	mh.n++
	mh.pq = append(mh.pq, k)
	mh.swim(mh.n)
}

// 上浮操作
// 对第k个元素进行上浮：如果比其父节点大，则替换两者位置，一直到堆顶
func (mh *MaxHeap) swim(k int) {
	if k > mh.n {
		return
	}
	for ; k > 1 && mh.pq[k] > mh.pq[k/2]; k /= 2 {
		mh.pq[k], mh.pq[k/2] = mh.pq[k/2], mh.pq[k]
	}
}

// 下沉
// 对第k个元素进行下沉，如果比其子节点小，则替换两者位置
func (mh *MaxHeap) sink(k int) {
	sub := k * 2
	if sub > mh.n {
		return
	}
	// 将父节点和子节点中较大的那个替换
	if sub+1 <= mh.n && mh.pq[sub+1] > mh.pq[sub] {
		sub++
	}
	if mh.pq[sub] < mh.pq[k] {
		mh.pq[sub], mh.pq[k] = mh.pq[k], mh.pq[sub]
		mh.sink(sub)
	}
}

// 最小堆
type MinHeap struct {
	pq []int // 优先队列 首个元素不存储值，对于父节点索引i，其子节点的索引分别为 2*i、2*i+1
	n  int   // 元素个数，则 n == len(pq)-1
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		pq: []int{0},
		n:  0,
	}
}

func (mh *MinHeap) isEmpty() bool {
	return mh.n == 0
}

// 删除最小元素：删除堆顶元素，并将末尾元素放到堆顶，对堆顶元素进行下沉
func (mh *MinHeap) delMin() int {
	if mh.isEmpty() {
		return 0
	}
	min := mh.pq[1]
	mh.pq[1] = mh.pq[mh.n]
	mh.n--
	mh.pq = mh.pq[:mh.n+1]
	mh.sink(1)
	return min
}

// 插入元素，将元素放到末尾，然后进行上浮
func (mh *MinHeap) insert(k int) {
	mh.n++
	mh.pq = append(mh.pq, k)
	mh.swim(mh.n)
}

// 上浮操作
// 对第k个元素进行上浮：如果比其父节点小，则替换两者位置，一直到堆顶
func (mh *MinHeap) swim(k int) {
	if k > mh.n {
		return
	}
	for ; k > 1 && mh.pq[k] < mh.pq[k/2]; k /= 2 {
		mh.pq[k], mh.pq[k/2] = mh.pq[k/2], mh.pq[k]
	}
}

// 下沉
// 对第k个元素进行下沉，如果比其子节点大，则替换两者位置
func (mh *MinHeap) sink(k int) {
	sub := k * 2
	if sub > mh.n {
		return
	}
	// 将父节点和子节点中较小的那个替换
	if sub+1 <= mh.n && mh.pq[sub+1] < mh.pq[sub] {
		sub++
	}
	if mh.pq[sub] < mh.pq[k] {
		mh.pq[sub], mh.pq[k] = mh.pq[k], mh.pq[sub]
		mh.sink(sub)
	}
}

// 有序矩阵的 Kth Element
// 给你一个n x n矩阵matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
func kthSmallest(matrix [][]int, k int) int {
	// 方法二：已知矩阵中的左上角是最小值，右上角是最大值，我们只要每次都拿当前矩阵的最小值，第k个值即为我们要找的值
	// 首次可以直接拿左上角的元素，第二次拿左上角相邻的两个元素中的最小值，第三次就不确定了，我们需要维护一个“候选列表”，从候选列表中获取最小值
	// 在第二次拿掉元素后，如果拿掉的是右上的元素，就需要将其右边的元素存入候选列表；如果拿掉的是左下的元素，就需要将其下方的元素存入候选列表
	// 这种每次都需要移出最小值并插入新值，适用于最小堆。
	// 堆中的元素需要维护其方向，下一个移入的元素是在其下边还是右边。但只要我们先把最左边的一列放入“候选列表”，则每次只需要移入右边的元素即可（
	// 如果先把最上边的一行放入候选列表，那么每次只需要移入下边的元素）
	n := len(matrix)
	heap := NewMinHeapKth()
	for i := 0; i < n; i++ {
		heap.insert(matrix[i][0], i, 0)
	}
	for i := 0; i < k-1; i++ {
		min := heap.delMin()
		x := min[2]
		if x < n-1 {
			y := min[1]
			heap.insert(matrix[y][x+1], y, x+1)
		}
	}
	//return heap.delMin()[0]
	a := heap.delMin()
	return a[0]
}

type MinHeapKth struct {
	pq [][3]int // 每个元素的三个值分别对应 值、纵坐标、横坐标
	n  int      // 元素个数，则 n == len(pq)-1
}

func NewMinHeapKth() *MinHeapKth {
	return &MinHeapKth{
		pq: [][3]int{{0, 0, 0}},
		n:  0,
	}
}

func (mh *MinHeapKth) isEmpty() bool {
	return mh.n == 0
}

// 删除最小元素：删除堆顶元素，并将末尾元素放到堆顶，对堆顶元素进行下沉
func (mh *MinHeapKth) delMin() [3]int {
	if mh.isEmpty() {
		return [3]int{}
	}
	min := mh.pq[1]
	mh.pq[1] = mh.pq[mh.n]
	mh.n--
	mh.pq = mh.pq[:mh.n+1]
	mh.sink(1)
	return min
}

// 插入元素，将元素放到末尾，然后进行上浮
func (mh *MinHeapKth) insert(k, y, x int) {
	mh.n++
	mh.pq = append(mh.pq, [3]int{k, y, x})
	mh.swim(mh.n)
}

// 上浮操作
// 对第k个元素进行上浮：如果比其父节点小，则替换两者位置，一直到堆顶
func (mh *MinHeapKth) swim(k int) {
	if k > mh.n {
		return
	}
	for ; k > 1 && mh.pq[k][0] < mh.pq[k/2][0]; k /= 2 {
		mh.pq[k], mh.pq[k/2] = mh.pq[k/2], mh.pq[k]
	}
}

// 下沉
// 对第k个元素进行下沉，如果比其子节点大，则替换两者位置
func (mh *MinHeapKth) sink(k int) {
	sub := k * 2
	if sub > mh.n {
		return
	}
	// 将父节点和子节点中较小的那个替换
	if sub+1 <= mh.n && mh.pq[sub+1][0] < mh.pq[sub][0] {
		sub++
	}
	if mh.pq[sub][0] < mh.pq[k][0] {
		mh.pq[sub], mh.pq[k] = mh.pq[k], mh.pq[sub]
		mh.sink(sub)
	}
}
