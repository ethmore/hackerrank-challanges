package dayone

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	a := strings.Split(s, ":")
	cycle := a[2][2:]
	hourInt, _ := strconv.Atoi(a[0])

	if cycle == "PM" && hourInt == 12 {

	} else if cycle == "PM" {
		hourInt += 12
		a[0] = strconv.Itoa(hourInt)
	} else if cycle == "AM" && hourInt == 12 {
		a[0] = "00"
	}

	sec := a[2][:2]
	x := fmt.Sprintf("%s:%s:%s", a[0], a[1], sec)
	return x
}
