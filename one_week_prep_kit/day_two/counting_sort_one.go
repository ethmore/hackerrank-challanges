package daytwo

func countingSort(arr []int32) []int32 {
	newArr := make([]int32, 100)

	for _, j := range arr {
		newArr[j]++
	}

	return newArr
}
