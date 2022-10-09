package arrays

import "fmt"

func minimumBribes(q []int32) {
	var bribes int32

	for i := range q {
		if q[i]-(int32(i)+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}
	}

	for m := len(q) - 1 - 2; m >= 0; m-- {
		for n := m; n <= m+2; n++ {
			if n < len(q)-1 && q[n] > q[n+1] {
				tmp := q[n]
				q[n] = q[n+1]
				q[n+1] = tmp
				bribes++
			}
		}
	}

	fmt.Println(bribes)
}
