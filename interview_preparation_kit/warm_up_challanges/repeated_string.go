package warmup

func repeatedString(s string, n int64) int64 {
	strLen := len(s)
	var aCount int64 = 0

	for _, c := range s {
		if c == 'a' {
			aCount++
		}
	}

	mul := n / int64(strLen)
	aCount = aCount * mul

	missigLetterCount := n - mul*int64(strLen)
	var j int64 = 0
	for j < missigLetterCount {
		if s[j] == 'a' {
			aCount++
		}
		j++
	}

	return aCount
}
