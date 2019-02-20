package solution

import (
	"fmt"
)

var hours = map[int][]string{
	0: []string{"0"},
	1: []string{"1", "2", "4", "8"},
	2: []string{"3", "5", "6", "9", "10"},
	3: []string{"7", "11"},
}

var minutes = map[int][]string{
	0: []string{"00"},
	1: []string{"01", "02", "04", "08", "16", "32"},
	2: []string{"03", "05", "06", "09", "10", "12", "17", "18", "20", "24", "33", "34", "36", "40", "48"},
	3: []string{"07", "11", "13", "14", "19", "21", "22", "25", "26", "28", "35", "37", "38", "41", "42", "44", "49", "50", "52", "56"},
	4: []string{"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58"},
	5: []string{"31", "47", "55", "59"},
}

func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	for i := 0; i <= num && i < len(hours); i++ {
		if num-i >= len(minutes) {
			continue
		}

		for _, h := range hours[i] {
			for _, m := range minutes[num-i] {
				res = append(res, fmt.Sprintf("%s:%s", h, m))
			}
		}
	}

	return res
}

func catOneCount(num int) int {
	var count int
	for num > 0 {
		if num&1 == 1 {
			count += 1
		}
		num >>= 1
	}
	return count
}
