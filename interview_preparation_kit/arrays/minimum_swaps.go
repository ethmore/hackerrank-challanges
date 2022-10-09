package arrays

import "math"

func minimumSwaps(arr []int32) int32 {
	indexOfBiggestDifference := int32(0)
	indexOfSecondBiggestDifference := int32(0)
	biggest := int32(0)
	secondBiggest := int32(0)
	swap := false
	swapCount := int32(0)

	for i, j := range arr {
		var absDifference int32 = int32(math.Abs(float64(j - int32((i + 1)))))
		if absDifference > indexOfBiggestDifference {
			indexOfBiggestDifference = absDifference
			biggest = arr[indexOfBiggestDifference]
			swap = true
		} else if absDifference > indexOfSecondBiggestDifference {
			indexOfSecondBiggestDifference = absDifference
			secondBiggest = arr[indexOfSecondBiggestDifference]
			swap = true
		}
		if swap {
			arr[indexOfBiggestDifference] = secondBiggest
			arr[indexOfSecondBiggestDifference] = biggest
			swapCount++
		}
	}
	return swapCount
}
