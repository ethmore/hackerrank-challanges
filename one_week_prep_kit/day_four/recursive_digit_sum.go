package dayfour

import "strconv"

func superDigit(n string, k int32) int32 {
	var sum int64

	for i := range n {
		sum = sum + int64(((int32(n[i]))-int32('0'))*int32(k))
	}
	if sum > 9 {
		sumStr := strconv.FormatInt(sum, 10)
		return superDigit(sumStr, 1)
	}

	return int32(sum)
}
