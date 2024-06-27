package tasks

func countMaxNumber(arr []int32) int32 {
	var max, count int32

	for _, num := range arr {
		if num > max {
			count = 1
			max = num
			continue
		}

		if num == max {
			count += 1
		}
	}

	return count
}
