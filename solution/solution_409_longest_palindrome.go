package solution

func LongestPalindrome(s string) int {
	charMap := make(map[int32]int, 0)
	for _, ch := range s {
		if _, ok := charMap[ch]; ok {
			charMap[ch] = charMap[ch] + 1
		} else {
			charMap[ch] = 1
		}
	}
	hasOdd := false
	maxLen := 0
	for _, count := range charMap {
		if !hasOdd && count&1 == 1 {
			hasOdd = true
		}
		maxLen += (count >> 1)
	}
	maxLen *= 2
	if hasOdd {
		maxLen++
	}
	return maxLen
}
