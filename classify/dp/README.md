# 动态规划

动态规划常用于解决“最优解”问题，其本质上仍是“穷举”，但是通过对“重叠子问题”的结果备份实现了“剪枝”。
## 关键点
1. 找到重叠子问题
2. 列出状态转移方程

## 问题划分
### 股票买卖
- 121. [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)（简单）

- 122. [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)（简单）

- 123. [买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)（困难）

- 188. [买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)（困难）

- 309. [最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)（中等）

- 714. [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)（中等）