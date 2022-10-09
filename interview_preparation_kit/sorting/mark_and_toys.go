package sorting

import "sort"

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	toys := 0
	for i := range prices {
		if k > 0 && k >= prices[i] {
			k = k - prices[i]
			toys++
		} else {
			return int32(toys)
		}
	}

	return int32(toys)
}
