package greedyalgorithms

import "sort"

func getMinimumCost(k int32, c []int32) int32 {
	sort.SliceStable(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	total := int32(0)

	for i := range c {
		total += (int32(i)/k + 1) * c[i]
	}
	return total
}
