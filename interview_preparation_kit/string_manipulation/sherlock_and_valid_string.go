package stringmanipulation

import "math"

func getMinMax(m map[int]int) (int, int) {
	maxVal := 0
	minVal := math.MaxInt64
	for i := range m {
		if maxVal < i {
			maxVal = i
		}
		if minVal > i {
			minVal = i
		}
	}

	return minVal, maxVal
}

func isValid(s string) string {
	runes := []rune(s)
	occur := make(map[rune]int)
	occurOfOccur := make(map[int]int)

	for _, val := range runes {
		occur[val]++
	}
	for _, val := range occur {
		occurOfOccur[val]++
	}

	if len(occurOfOccur) == 1 {
		return "YES"
	} else if len(occurOfOccur) == 2 {
		minVal, maxVal := getMinMax(occurOfOccur)

		if maxVal-minVal == 1 && occurOfOccur[maxVal] == 1 {
			return "YES"
		} else if minVal == 1 && occurOfOccur[minVal] == 1 {
			return "YES"
		}
	}
	return "NO"
}
