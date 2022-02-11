package stack_queue

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
