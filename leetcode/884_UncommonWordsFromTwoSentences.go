package leetcode

/*
- 884.Uncommon Words From Two Sentences
- Description:
We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.

Example 1:
Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]

Example 2:
Input: A = "apple apple", B = "banana"
Output: ["banana"]

Note:
0 <= A.length <= 200
0 <= B.length <= 200
A and B both contain only spaces and lowercase letters.
*/
import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	wordMaps := make(map[string]int)
	returnList := make([]string, 0)
	for _, word := range strings.Split(A, " ") {
		wordMaps[word]++
	}
	for _, word := range strings.Split(B, " ") {
		wordMaps[word]++
	}
	for word, cnt := range wordMaps {
		if cnt == 1 {
			returnList = append(returnList, word)
		}
	}
	return returnList
}
