package sorting

import "fmt"

func countSwaps(a []int32) {
	swaps := 0
	for range a {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swaps++
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps.", swaps)
	fmt.Println()
	fmt.Printf("First Element: %d", a[0])
	fmt.Println()
	fmt.Printf("Last Element: %d", a[len(a)-1])
}
