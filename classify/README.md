# 分类题

## 资料
- [LeetCode分类参考](https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E7%9B%AE%E5%BD%95.md)
- [告别动态规划，连刷 40 道题，我总结了这些套路，看不懂你打我](https://zhuanlan.zhihu.com/p/91582909)
- [代码随想录](https://github.com/youngyangyang04/leetcode-master)
- [OI Wiki](https://oi-wiki.org/)

对于搜索问题：
- 如果是在网格中找到最短路径，那么就是广度优先
- 如果是在网格中找到“可行路径“，那么就是深度优先
- 如果是找到最优解，那么就是动态规划（存在限定条件：只能向右和向下）

### 动态规划思路
1. 定义数组元素的含义，即找到dp[i][j]
2. 找出关系数组元素间的关系式, 即dp[i][j]与dp[i-1][j-1]或dp[i-1][j]或dp[i][j-1]之间的关联
3. 初始化数组，dp[i][0]，dp[0][j]往往需要初始化，因为在计算时总要用到dp[i-1][j-1]

01背包
- [目标和](https://leetcode-cn.com/problems/target-sum/) ：转化问题以后为0-1背包方案数问题。
- [分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/) ：转化后为0-1背包可行性问题。
- [最后一块石头的重量 II](https://leetcode-cn.com/problems/last-stone-weight-ii/) ：转化后为0-1背包最小值问题。
完全背包
- [零钱兑换](https://leetcode-cn.com/problems/coin-change/) ：完全背包最小值
- [完全平方数](https://leetcode-cn.com/problems/perfect-squares/) ：完全背包最小值
- [零钱兑换 II](https://leetcode-cn.com/problems/coin-change-2/) ：完全背包方案数
- [组合总和 Ⅳ](https://leetcode-cn.com/problems/combination-sum-iv/) ：考虑物品顺序的完全背包方案数。每个物品可以重复拿，有几种装满背包的方案？
多维背包
- [01 字符构成最多的字符串](https://leetcode-cn.com/problems/ones-and-zeroes/) ：多维费用的 0-1 背包最大值，两个背包大小：0和1的数量
- [盈利计划](https://leetcode-cn.com/problems/profitable-schemes/) ：多维费用的 0-1 背包最大值
分组背包
- [掷骰子的N种方法](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/) ：每一组是一个骰子，每个骰子只能拿一个体积为1到6的物品


## 解题步骤
1. 明确问题
2. 简化问题（可选）
    - 如目标和问题findTargetSumWays
3. 问题划分
4. 处理特殊情况（可选）
5. 实现

## TODO
1. 动态规划：两层循环的先后顺序，以及空闲优化后的内存循环逆序问题（change、change2）
2. 动态规划：combinationSum4、coinChange、change之间的区别总结
4. [morrois遍历](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/solution/ba-er-cha-sou-suo-shu-zhuan-huan-wei-lei-jia-sh-14/)