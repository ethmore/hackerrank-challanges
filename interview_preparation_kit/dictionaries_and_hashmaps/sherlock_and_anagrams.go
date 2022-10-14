package dictionariesandhashmaps

import "sort"

func sherlockAndAnagrams(s string) int32 {
	mp := make(map[string]int32)
	subS := ""
	anagrams := int32(0)
	for subLen := 0; subLen < len(s); subLen++ {
		subS = ""
		for subStr := 0; subStr < len(s)-subLen; subStr++ {
			subS += string(s[subLen+subStr])
			subS = sortString(subS)
			mp[subS]++
		}
	}

	for _, val := range mp {
		anagrams += (val * (val - 1) / 2)
	}
	return anagrams
}

func sortString(s string) string {
	runes := []rune(s)
	sort.SliceStable(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
