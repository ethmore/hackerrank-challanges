package warmup

func sockMerchant(n int32, ar []int32) int32 {
	var count int32

	m := make(map[int32]bool)

	for _, v := range ar {
		if m[v] {
			count++
			m[v] = false
		} else {
			m[v] = true
		}
	}

	return count
}
