package search

import "fmt"

func whatFlavors(cost []int32, money int32) {
	saved := make(map[int32]int)
	for i, val := range cost {
		if idx, f := saved[money-val]; f { // if remaining money in saved
			fmt.Println(idx, i+1)
			return
		}
		saved[val] += i + 1
	}
}
