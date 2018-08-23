package leetcode

/*
- 344.Reverse String
- Description:
Write a function that takes a string as input and returns the string reversed.

Example 1:
Input: "hello"
Output: "olleh"

Example 2:
Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"
*/

func reverseString(s string) string {
	reverseStr := []rune(s)
	length := len(reverseStr)
	for i := 0; i < length/2; i++ {
		reverseStr[i], reverseStr[length-1-i] = reverseStr[length-1-i], reverseStr[i]
	}
	return string(reverseStr)
}
