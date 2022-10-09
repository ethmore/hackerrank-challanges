package warmup

func jumpingOnClouds(c []int32) int32 {
	length := len(c)
	var jumpCount int32 = 0

	for i := 0; i < length-1; {
		if i+2 == length || c[i+2] == 1 {
			jumpCount++
			i++
		} else {
			jumpCount++
			i += 2
		}
	}

	return jumpCount
}
