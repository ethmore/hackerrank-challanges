package dayone

import "fmt"

func plusMinus(arr []int32) {
	positives := float32(0)
	negatives := float32(0)
	zeros := float32(0)
	length := float32(len(arr))

	for _, j := range arr {
		if j > 0 {
			positives++
		} else if j < 0 {
			negatives++
		} else {
			zeros++
		}
	}
	fmt.Println(positives / length)
	fmt.Println(negatives / length)
	fmt.Println(zeros / length)
}
