package daytwo

import "math"

func diagonalDifference(arr [][]int32) int32 {
	length := len(arr)
	leftToR := int64(0)
	rightToL := int64(0)

	for i := 0; i < length; i++ {
		leftToR += int64(arr[i][i])
		rightToL += int64(arr[i][length-(i+1)])
	}

	return int32(math.Abs(float64(leftToR - rightToL)))
}
