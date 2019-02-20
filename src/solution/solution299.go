package solution

import "fmt"

func getHint(secret string, guess string) string {
	var countA, countB int
	length := len(secret)
	for i := 0; i < length; i++ {
		if guess[i] == secret[i] {
			countA = countA + 1
		} else {
			for j := 0; j < length; j++ {
				if guess[i] == secret[j] {
					countB = countB + 1
				}
			}
		}
	}

	return fmt.Sprintf("%vA%vB", countA, countB)
}
