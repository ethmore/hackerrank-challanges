package arrays

func hourglassSum(arr [][]int32) int32 {
	var maxSum int32

	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			firstLine := arr[a][b] + arr[a][b+1] + arr[a][b+2]
			secondLine := arr[a+1][b+1]
			thirdLine := arr[a+2][b] + arr[a+2][b+1] + arr[a+2][b+2]
			sum := firstLine + secondLine + thirdLine
			if a == 0 && b == 0 {
				maxSum = sum
			} else {
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}

	return maxSum
}
