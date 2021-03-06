package leetcode

/*
- 70. Climbing Stairs
- Description
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

- Example1:
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

- Example2:
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	cache := make([]int, n+1)
	cache[1] = 1
	cache[2] = 2
	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}
