# [1029. Two City Scheduling](https://leetcode.com/problems/two-city-scheduling/)

## 2019/06/04

### 题目 💗[easy]

There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

公司计划采访 2N 个人.飞机费用是 i-th 个人 去城市 A,花费`costs[i][0]`,去城市 B 花费`costs[i][1]`.

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.

返回最小花费

Example 1:

```bash
Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.
```

Note:

1. 1 <= costs.length <= 100
2. It is guaranteed that costs.length is even.
3. 1 <= costs[i][0], costs[i][1] <= 1000
