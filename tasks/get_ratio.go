package tasks

import "fmt"

func getRatiosFromArray(arr []int32) {
	var positives, negatives, zeros int32

	for _, num := range arr {
		switch {
		case num > 0:
			positives++
		case num < 0:
			negatives++
		default:
			zeros++
		}
	}

	total := float32(len(arr))
	fmt.Printf("%.6f\n", float32(positives)/total)
	fmt.Printf("%.6f\n", float32(negatives)/total)
	fmt.Printf("%.6f\n", float32(zeros)/total)
}
