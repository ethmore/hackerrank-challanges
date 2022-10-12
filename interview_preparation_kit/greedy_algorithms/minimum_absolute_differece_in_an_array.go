package greedyalgorithms

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	var minDiff int32 = math.MaxInt32
	var absDiff int32 = 0
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	for i := range arr {
		if i+1 < len(arr) {
			absDiff = int32(math.Abs(float64(arr[i+1]) - float64(arr[i])))

			if absDiff < minDiff {
				minDiff = absDiff
			}
		}
	}
	return minDiff
}
