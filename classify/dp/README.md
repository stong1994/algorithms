# 动态规划

动态规划常用于解决“最优解”问题，其本质上仍是“穷举”，但是通过对“重叠子问题”的结果备份实现了“剪枝”。
## 关键点
1. 找到重叠子问题
2. 列出状态转移方程
3. 计算最优解的值，通常采用自底向上的方式先求得子问题的最优解在求得原问题的最优解
最优子结构性质：问题的最优解由相关子问题的最优解组合而成，而这些子问题可以独立求解



## 问题划分
### 股票买卖
- LeetCode121. [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)（简单）
- LeetCode122. [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)（简单）
- LeetCode123. [买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)（困难）
- LeetCode188. [买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)（困难）
- LeetCode309. [最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)（中等）
- LeetCode714. [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)（中等）

- 切割钢铁——《算法导论》