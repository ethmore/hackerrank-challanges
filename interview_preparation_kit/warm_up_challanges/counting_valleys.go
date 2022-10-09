package warmup

func countingValleys(steps int32, path string) int32 {
	var altitude int32 = 0
	var oldAltitude int32 = 0
	var valleyCount int32 = 0

	for _, elevation := range path {
		if elevation == 'U' {
			altitude++
		} else {
			altitude--
		}

		if altitude < 0 && oldAltitude == 0 {
			valleyCount++
		}
		oldAltitude = altitude
	}

	return valleyCount
}
