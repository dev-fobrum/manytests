package tasks

import (
	"fmt"
	"testing"
)

type CountMaxNumberProps struct {
	arr         []int32
	expect      int32
	expectError *bool
}

func TestCountMaxNumber(t *testing.T) {
	tests := []CountMaxNumberProps{
		{arr: []int32{3, 2, 1, 3}, expect: 2},
		{arr: []int32{3, 2, 1, 3, 5, 4, 2, 1, 2}, expect: 1},
		{arr: []int32{3, 3, 3, 3}, expect: 4},
		{arr: []int32{10, 12, 15, 13, 100, 100, 72, 55}, expect: 2},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("For %d expect %d", test.arr, test.expect)

		t.Run(testname, func(t *testing.T) {
			result := countMaxNumber(test.arr)

			if result != test.expect {
				t.Errorf("countMaxNumber(%v) = %v; want %v", test.arr, result, test.expect)
			}
		})
	}
}
