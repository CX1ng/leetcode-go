package leetcode

/*
- 746.Min Cost Climbing Stairs
- Description:
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

- Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

- Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
*/
import (
	. "github.com/CX1ng/leetcode-go/utils"
)

func minCostClimbingStairs(cost []int) int {
	length := len(cost) + 1

	if length == 0 || length == 1 {
		return 0
	}

	cache := make([]int, length)
	for i := 2; i < length; i++ {
		cache[i] = Min(cache[i-1]+cost[i-1], cache[i-2]+cost[i-2])
	}
	return cache[length-1]
}
