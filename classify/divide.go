package classify

import (
	"fmt"
	"strconv"
)

// 分治算法

// 为运算表达式设计优先级
// 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。
// 有效的运算符号包含 +,-以及*。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/different-ways-to-add-parentheses

func diffWaysToCompute(expression string) []int {
	return allRst(expression, 0, len(expression)-1)
}

func allRst(expression string, lo int, hi int) []int {
	var symbol = map[int32]struct{}{'+': {}, '-': {}, '*': {}}
	result := make([]int, 0)
	for i := lo; i <= hi; i++ {
		if _, ok := symbol[int32(expression[i])]; ok {
			leftNums := allRst(expression, lo, i-1)
			rightNums := allRst(expression, i+1, hi)
			allNums := calc(leftNums, rightNums, expression[i])
			for _, v := range allNums {
				result = append(result, v)
			}
		}
	}
	if len(result) == 0 {
		rst, _ := strconv.Atoi(expression[lo : hi+1])
		return []int{rst}
	}
	return result
}

func calc(nums1 []int, nums2 []int, sym byte) []int {
	var result []int
	switch sym {
	case '+':
		for _, v := range nums1 {
			for _, b := range nums2 {
				result = append(result, v+b)
			}
		}
	case '-':
		for _, v := range nums1 {
			for _, b := range nums2 {
				result = append(result, v-b)
			}
		}
	case '*':
		for _, v := range nums1 {
			for _, b := range nums2 {
				result = append(result, v*b)
			}
		}
	default:
		panic("invalid sym")
	}
	return result
}

// 不同的二叉搜索树
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func printTreeNode(idx int, pre string, node *TreeNode) {
	if node == nil {
		fmt.Printf("%*s nil", idx*4, pre)
		return
	}
	fmt.Printf("%*s", idx*4, "")
	fmt.Println("root:", node.Val)
	idx++
	printTreeNode(idx, "left:", node.Left)
	printTreeNode(idx, " right:", node.Right)
	println("")
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return getNodes(1, n)
}

func getNodes(lo int, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}
	if lo == hi {
		return []*TreeNode{{Val: lo}}
	}
	var result []*TreeNode
	for i := lo; i <= hi; i++ {
		lefts := getNodes(lo, i-1)
		rights := getNodes(i+1, hi)
		for _, l := range lefts {
			for _, r := range rights {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return result
}
