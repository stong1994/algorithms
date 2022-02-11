package stack_queue

import (
	"math"
)

// 用队列实现栈
// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
//实现 MyStack 类：
//	void push(int x) 将元素 x 压入栈顶。
//	int pop() 移除并返回栈顶元素。
//	int top() 返回栈顶元素。
//	boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
// 注意：
// 你只能使用队列的基本操作 —— 也就是push to back、peek/pop from front、size 和is empty这些操作。
// 你所使用的语言也许不支持队列。你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列, 只要是标准的队列操作即可。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-stack-using-queues
type MyStack struct {
	list []int
}

func ConstructorStack() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyStack) Pop() int {
	v := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return v
}

func (this *MyStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.list) == 0
}

// 用栈实现队列
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
// 实现 MyQueue 类：
//	void push(int x) 将元素 x 推到队列的末尾
//	int pop() 从队列的开头移除并返回元素
//	int peek() 返回队列开头的元素
//	boolean empty() 如果队列为空，返回 true ；否则，返回 false
// 说明：
//	你 只能 使用标准的栈操作 —— 也就是只有push to top,peek/pop from top,size, 和is empty操作是合法的。
//	你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	v := this.list[0]
	this.list = this.list[1:]
	return v
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.list) == 0
}

// 最小值栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop()—— 删除栈顶的元素。
// top()—— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/min-stack
type MinStack struct {
	// 区别于普通的栈，增加了获取栈中最小元素的功能，用另一个辅助栈来实现即可
	list       []int
	assistList []int // 辅助栈中的元素与list中的元素一一对应，随list pop一起pop，区别在于辅助栈中对应的元素为已有元素的最小值
}

func ConstructorMinStack() MinStack {
	return MinStack{
		list:       nil,
		assistList: []int{math.MaxInt32}, // 初始化一个值有利于后边的push操作
	}
}

func (this *MinStack) Push(val int) {
	this.list = append(this.list, val)
	small := this.assistList[len(this.assistList)-1]
	this.assistList = append(this.assistList, min(small, val))
}

func (this *MinStack) Pop() {
	this.list = this.list[:len(this.list)-1]
	this.assistList = this.assistList[:len(this.assistList)-1]
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	return this.assistList[len(this.assistList)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 用栈实现括号匹配
// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
//	左括号必须用相同类型的右括号闭合。
//	左括号必须以正确的顺序闭合。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-parentheses
func isValid(s string) bool {
	isMatch := func(a, b int32) bool {
		return (a == '(' && b == ')') || (a == '{' && b == '}') || (a == '[' && b == ']')
	}
	var list []int32
	for _, v := range s {
		L := len(list)
		if L == 0 {
			list = append(list, v)
		} else if isMatch(list[L-1], v) {
			list = list[:L-1]
		} else {
			list = append(list, v)
		}
	}
	return len(list) == 0
}

// 数组中元素与下一个比它大的元素之间的距离
// 给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指在第 i 天之后，才会有更高的温度。
// 如果气温在这之后都不会升高，请在该位置用0 来代替。
// 提示：
// 1 <= temperatures.length <= 105
// 30 <= temperatures[i] <= 100
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/daily-temperatures
func dailyTemperatures(temperatures []int) []int {
	// 靠前的元素对应的结果后出，符合“栈”
	// 构建一个递减栈，顺序遍历temperatures，如果当前元素大于栈顶元素，则将栈顶元素pop，并对比两者的索引，其差值为目标值。
	// 如果当前元素小于栈顶元素，则将其入栈
	var stack [][2]int
	L := len(temperatures)
	result := make([]int, L)
	for i := 0; i < L; i++ {
		cur := temperatures[i]
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if cur > top[1] {
				result[top[0]] = i - top[0]
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, [2]int{i, cur})
	}
	return result
}

// 循环数组中比当前元素大的下一个元素
// 给定一个循环数组nums（nums[nums.length - 1]的下一个元素是nums[0]），返回nums中每个元素的 下一个更大元素 。
//数字 x的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/next-greater-element-ii
func nextGreaterElements(nums []int) []int {
	// 遍历两次，构建递减栈。用当前元素与栈顶元素比较，如果当前元素大于栈顶元素，则pop栈顶元素；否则将当前元素入栈
	//result := make([]int, len(nums))
	//for i := range result {
	//	result[i] = -1
	//}
	//stack := make([][2]int, 0)
	//L := len(nums)
	//for j := 0; j <= 1; j++ {
	//	for i := 0; i < len(nums); i++ {
	//		cur := nums[i]
	//		for len(stack) > 0 {
	//			top := stack[len(stack)-1]
	//			if top[1] >= cur {
	//				break
	//			}
	//			idx := top[0]
	//			if idx >= L {
	//				idx -= L
	//			}
	//			result[idx] = cur
	//			stack = stack[:len(stack)-1]
	//		}
	//		stack = append(stack, [2]int{L*j + i, cur})
	//	}
	//}
	//return result

	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	stack := make([][2]int, 0)
	L := len(nums)
	for i := 0; i < 2*L-1; i++ {
		cur := nums[i%L]
		for ; len(stack) > 0; stack = stack[:len(stack)-1] {
			top := stack[len(stack)-1]
			if cur <= top[1] {
				break
			}
			result[top[0]] = cur
		}
		stack = append(stack, [2]int{i % L, cur})
	}
	return result
}
