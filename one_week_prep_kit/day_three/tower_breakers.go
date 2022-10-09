package daythree

func towerBreakers(n int32, m int32) int32 {
	if n%2 == 0 {
		return 2
	} else {
		if m == 1 {
			return 2
		}
	}

	return 1
}
