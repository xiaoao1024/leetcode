package solution

func IsPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 == 0 {
			n = n / 3
		} else {
			return false
		}
	}
	return true
}
