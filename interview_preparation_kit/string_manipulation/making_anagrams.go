package stringmanipulation

import "math"

func makeAnagram(a string, b string) int32 {
	var delCnt int32 = 0
	aMap := make(map[rune]int32)
	for _, char := range a {
		aMap[char]++
	}
	for _, char := range b {
		_, found := aMap[char]
		if !found {
			aMap[96]++
		} else {
			aMap[char]--
		}
	}

	for _, val := range aMap {
		delCnt += int32(math.Abs(float64(val)))
	}

	return delCnt
}
