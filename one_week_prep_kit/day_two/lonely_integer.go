package daytwo

func lonelyinteger(a []int32) int32 {
	lone := int32(0)
	m := make(map[int32]int)

	for _, j := range a {
		m[j]++
	}

	for k, v := range m {
		if v != 2 {
			lone = int32(k)
		}
	}

	return lone
}
