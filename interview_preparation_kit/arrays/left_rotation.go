package arrays

func rotLeft(a []int32, d int32) []int32 {
	var i int32

	for i = 0; i < d; i++ {
		a = append(a[1:], a[0])
	}
	return a
}
