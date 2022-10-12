package greedyalgorithms

import "sort"

func luckBalance(k int32, contests [][]int32) int32 {
	var finalLuck int32 = 0
	sort.SliceStable(contests, func(i, j int) bool {
		return contests[i][0] > contests[j][0]
	})
	sort.SliceStable(contests, func(i, j int) bool {
		return contests[i][1] > contests[j][1]
	})

	for i := 0; i < len(contests); i++ {
		if int32(i) < k {
			finalLuck += contests[i][0]
		} else if contests[i][1] == 0 {
			finalLuck += contests[i][0]
		} else {
			finalLuck -= contests[i][0]
		}
	}

	return finalLuck
}
