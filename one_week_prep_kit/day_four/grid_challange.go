package dayfour

import "sort"

func gridChallenge(grid []string) string {
	var newArr [][]rune

	for _, j := range grid {
		s := []rune(j)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

		newArr = append(newArr, s)
	}

	for i := 0; i < len(newArr[0]); i++ {
		for k := 0; k < len(newArr)-1; k++ {
			if newArr[k][i] > newArr[k+1][i] {
				return "NO"
			}
		}
	}

	return "YES"
}
