package tasks

import "errors"

func compareBiggerThanInArrays(a []int32, b []int32) ([]int32, error) {
	if len(a) != len(b) {
		return []int32{0, 0}, errors.New("Arrays don't have the same lenght")
	}

	points := []int32{0, 0}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] > b[i] {
			points[0] += 1
		} else {
			points[1] += 1
		}
	}

	return points, nil
}
