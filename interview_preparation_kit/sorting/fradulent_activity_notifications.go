package sorting

import "math"

func activityNotifications(expenditure []int32, d int32) int32 {

	notify := int32(0)
	medianLeft, medianRight := (d-1)/2, int32(math.Ceil((float64(d)-1)/2))
	var mLeft, mRight int32
	var avgMed float64

	expenCount := make(map[int32]int32, 201)
	for i := d - 1; i >= 0; i-- {
		expenCount[expenditure[i]]++
	}

	for i := d; i < int32(len(expenditure)); i++ {
		for j, k := int32(0), int32(0); k <= medianLeft; j++ {
			k += expenCount[j]
			mLeft = j
		}

		for j, k := int32(0), int32(0); k <= medianRight; j++ {
			k += expenCount[j]
			mRight = j
		}
		avgMed = float64((mLeft + mRight)) / 2

		if float64(expenditure[i]) >= avgMed*2 {
			notify++
		}

		expenCount[expenditure[i-d]]--
		expenCount[expenditure[i]]++
	}
	return notify
}
