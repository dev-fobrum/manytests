package tasks

import (
	"bytes"
	"fmt"
	"os"
	"syscall"
	"testing"
)

type TestGetRatiosFromArrayProps struct {
	arr    []int32
	expect string
}

func TestGetRatiosFromArray(t *testing.T) {
	tests := []TestGetRatiosFromArrayProps{
		{arr: []int32{-4, 3, -9, 0, 4, 1}, expect: "0.500000\n0.333333\n0.166667\n"},
		{arr: []int32{-4, 3, -9, 0, 4, 1, 4, -8, 0, 5, -1}, expect: "0.454545\n0.363636\n0.181818\n"},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Testing %d", i)

		t.Run(testname, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			getRatiosFromArray(test.arr)

			w.Close()
			var buf bytes.Buffer
			buf.ReadFrom(r)
			actualOutput := buf.String()

			os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

			if actualOutput != test.expect {
				t.Errorf("expected %q but got %q", test.expect, actualOutput)
			}
		})
	}
}
