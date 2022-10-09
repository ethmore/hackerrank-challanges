package dayone

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	minVal := int64(0)
	maxVal := int64(0)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < 4; i++ {
		minVal = minVal + int64(arr[i])
	}
	for j := 4; j > 0; j-- {
		maxVal = maxVal + int64(arr[j])
	}

	fmt.Printf("%d %d\n", minVal, maxVal)
}
