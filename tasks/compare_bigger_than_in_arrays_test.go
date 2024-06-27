package tasks

import (
	"fmt"
	"testing"
)

func auxCheckIfArraysAreEqual(a []int32, b []int32) bool {
	equal := true

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			equal = false
			break
		}
	}

	return equal
}

type CompareBiggerThanInArraysProps struct {
	a           []int32
	b           []int32
	expect      []int32
	expectError bool
}

func TestCompareBiggerThanInArrays(t *testing.T) {
	tests := []CompareBiggerThanInArraysProps{
		{a: []int32{1, 5, 7}, b: []int32{1, 3, 8}, expect: []int32{1, 1}},
		{a: []int32{1, 5, 7}, b: []int32{5, 10, 10}, expect: []int32{0, 3}},
		{a: []int32{51, 55, 57}, b: []int32{5, 10, 10}, expect: []int32{3, 0}},
		{a: []int32{1, 1, 1}, b: []int32{1, 1, 1}, expect: []int32{0, 0}},
		{a: []int32{1, 1, 1, 1}, b: []int32{1, 1, 1}, expect: []int32{0, 0}, expectError: true},
		{a: []int32{1, 1, 1}, b: []int32{1, 1}, expect: []int32{0, 0}, expectError: true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("For %d and %d expect %d", test.a, test.b, test.expect)

		t.Run(testname, func(t *testing.T) {
			result, err := compareBiggerThanInArrays(test.a, test.b)

			if test.expectError {
				if err == nil {
					t.Errorf("compareTriplets(%v, %v) expected error but got none", test.a, test.b)
				}
			} else {
				equal := auxCheckIfArraysAreEqual(result, test.expect)

				if !equal {
					t.Errorf("compareTriplets(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
				}
			}

		})
	}
}
