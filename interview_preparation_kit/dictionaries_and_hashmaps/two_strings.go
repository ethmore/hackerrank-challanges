package dictionariesandhashmaps

func twoStrings(s1 string, s2 string) string {
	s1Runes := []rune(s1)
	s2Runes := []rune(s2)
	commonMap := make(map[rune]bool)

	for _, j := range s1Runes {
		commonMap[j] = false
	}

	for _, m := range s2Runes {
		_, ok := commonMap[m]
		if ok {
			return "YES"
		}
	}

	return "NO"
}
