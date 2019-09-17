package solution

import (
	"fmt"
)

func FirstUniqChar(s string) int {
	myMap := make(map[int32]int, 0)
	for _, char := range s {
		if count, ok := myMap[char]; ok {
			myMap[char] = count + 1
		} else {
			myMap[char] = 1
		}
	}
	fmt.Println(myMap)
	for index, char := range s {
		if myMap[char] == 1 {
			return index
		}
	}
	return -1
}
