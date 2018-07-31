package leetcode
/*
- 198. House Robber
- Description:
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

- Example 1:
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.

- Example 2:
Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
 */
import (
	. "github.com/CX1ng/leetcode-go/utils"
)


func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	cache := make([]int, length)
	cache[0] = nums[0]
	cache[1] = nums[1]
	if nums[1] < nums[0] {
		cache[1] = cache[0]
	}
	max := cache[1]
	for i := 2;i < length;i++ { // max(cache[i-2] + nums[i], cache[i-1])
		cache[i] = Max(cache[i-2] + nums[i], cache[i-1])
		max = Max(cache[i], max)
	}
	return max
}