package solution

import (
	"fmt"
	"testing"
)

func TestIndexLongestPalindromeSubstring(t *testing.T) {
	var (
		s          = "abccbd"
		start, end int
	)

	start, end = indexLongestPalindromeSubstring(s, 2, 3)
	fmt.Printf("[%d, %d]\n", start, end)
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("abccbd"))
	fmt.Println(longestPalindrome("babad"))
}
