package solution

func isIsomorphic(s string, t string) bool {
	preIndexOfS := make(map[byte]int, 0)
	preIndexOfT := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc := t[i]
		if preIndexOfS[sc] != preIndexOfT[tc] {
			return false
		}

		preIndexOfS[sc] = i + 1
		preIndexOfT[tc] = i + 1
	}

	return true
}
