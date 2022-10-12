package sorting

import (
	"fmt"
	"sort"
)

func main() {
	var tests int
	var n string
	var s int
	type student struct {
		name  string
		score int
	}
	var students []student

	fmt.Scan(&tests)
	for i := 0; i < tests; i++ {
		fmt.Scan(&n)
		fmt.Scan(&s)
		stu := student{
			name:  n,
			score: s,
		}
		students = append(students, stu)
	}

	sort.SliceStable(students, func(i, j int) bool {
		return students[i].name < students[j].name
	})

	sort.SliceStable(students, func(i, j int) bool {
		return students[i].score > students[j].score
	})

	for _, stu := range students {
		fmt.Printf("%s %d", stu.name, stu.score)
		fmt.Println()
	}
}
