// 1768. Merge Strings Alternately
// Source: https://leetcode.com/problems/merge-strings-alternately

package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abc", "defg"))  // Output: "adbecfg"
	fmt.Println(mergeAlternately("abcd", "efgh")) // Output: "aebfcgdh"
}

func mergeAlternately(word1, word2 string) string {
	result := ""

	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}

		i++
	}

	return result
}
