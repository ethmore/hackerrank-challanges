package stringmanipulation

func alternatingCharacters(s string) int32 {
	runes := []rune(s)
	count := int32(0)

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			runes[i] = 0
			count++
		}
	}
	return count
}
