package arrays

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int32, n+1)
	maxVal := int64(-1)
	count := int64(0)

	for _, j := range queries {
		arr[j[0]-1] += j[2]
		arr[j[1]] -= j[2]
	}

	for _, k := range arr {
		count += int64(k)
		if int64(count) > maxVal {
			maxVal = count
		}
	}
	return maxVal
}
