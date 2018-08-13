package leetcode

/*
- 654.Maximum Binary Tree
- Description:
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

- Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
Note:
The size of the given array will be in the range [1,1000].
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	// Search Max Value
	maxIndex := 0
	max := nums[0]
	for i, num := range nums {
		if num > max {
			maxIndex = i
			max = num
		}
	}
	// construct Binary Tree
	node := &TreeNode{
		Val: max,
	}
	if maxIndex > 0 {
		node.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	}
	if maxIndex < length-1 {
		node.Right = constructMaximumBinaryTree(nums[maxIndex+1 : length])
	}
	return node
}
