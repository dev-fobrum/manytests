package tasks

import (
	"fmt"
	"testing"
)

type ConvertTimeProps struct {
	time        string
	expect      string
	expectError bool
}

func TestConvertTime(t *testing.T) {
	tests := []ConvertTimeProps{
		{time: "07:05:45PM", expect: "19:05:45"},
		{time: "12:40:22AM", expect: "00:40:22"},
		{time: "12:45:54PM", expect: "12:45:54"},
		{time: "007:05:45PM", expect: "Error", expectError: true},
		{time: "AA:05:45PM", expect: "Error", expectError: true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("For %v expect %v", test.time, test.expect)

		t.Run(testname, func(t *testing.T) {
			res, err := convertTimeAMPMToMilitary(test.time)

			if test.expectError {
				if err == nil {
					t.Errorf("convertTimeAMPMToMilitary(%s) expected error but got none", test.time)
				}
			} else {
				if res != test.expect {
					t.Errorf("convertTimeAMPMToMilitary(%s) = %s; want %s", test.time, res, test.expect)
				}
			}

		})
	}
}
