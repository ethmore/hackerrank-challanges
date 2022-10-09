package dictionariesandhashmaps

import "fmt"

func checkMagazine(magazine []string, note []string) {
	magMap := make(map[string]int)
	for _, k := range magazine {
		magMap[k]++
	}

	for _, j := range note {
		magMap[j]--
		if magMap[j] < 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
