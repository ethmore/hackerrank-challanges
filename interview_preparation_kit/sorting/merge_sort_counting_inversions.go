package sorting

func countInversions(arr []int32) int64 {
	_, c := mergeSort(arr)
	return c
}

func mergeSort(arr []int32) (arrRet []int32, swaps int64) {
	if len(arr) < 2 {
		return arr, 0
	}
	first, fSwaps := mergeSort(arr[:len(arr)/2])
	second, sSwaps := mergeSort(arr[len(arr)/2:])
	res, mSwaps := merge(first, second)
	return res, mSwaps + fSwaps + sSwaps
}

func merge(a []int32, b []int32) ([]int32, int64) {
	final := []int32{}
	i := 0
	j := 0
	inversionCnt := int64(0)
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
			inversionCnt += int64(len(a)) - int64(i)
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final, inversionCnt
}
