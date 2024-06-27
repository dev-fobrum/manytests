## Example README: Testing in Go

### Introduction

This README provides a guide on how to write and run tests in Go using the `go test` command. It includes examples to help you get started with testing your Go code effectively.

### Code Example: Comparing Arrays

#### `compare_bigger_than_in_arrays.go`

```go
package tasks

import "errors"

func compareBiggerThanInArrays(a []int32, b []int32) ([]int32, error) {
	if len(a) != len(b) {
		return []int32{0, 0}, errors.New("Arrays don't have the same length")
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
```

#### `compare_bigger_than_in_arrays_test.go`

```go
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
		testname := fmt.Sprintf("For %v and %v expect %v", test.a, test.b, test.expect)

		t.Run(testname, func(t *testing.T) {
			result, err := compareBiggerThanInArrays(test.a, test.b)

			if test.expectError {
				if err == nil {
					t.Errorf("compareBiggerThanInArrays(%v, %v) expected error but got none", test.a, test.b)
				}
			} else {
				equal := auxCheckIfArraysAreEqual(result, test.expect)

				if !equal {
					t.Errorf("compareBiggerThanInArrays(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
				}
			}
		})
	}
}
```

### Running Tests

To run tests in Go, you use the `go test` command. Here are some common ways to run tests:

- **Running all tests (`-v` for verbose output)**:
  
  ```bash
  go test -v ./tasks
  ```

  This command runs all the tests in the current package (`tasks` in this example) and provides verbose output, showing the names of each test being executed and their results.

- **Running specific tests by name (`-run` with test name and `-v` for verbose output)**:

  ```bash
  go test -run TestCompareBiggerThanInArrays -v ./tasks
  ```

  This command runs only the test named `TestCompareBiggerThanInArrays` and provides verbose output. Replace `TestCompareBiggerThanInArrays` with the name of the specific test you want to run.

### Conclusion

This README provides a basic overview of writing and running tests in Go using a practical example. By following these practices, you can ensure that your Go code is thoroughly tested and reliable.

For more information on testing in Go, refer to the official documentation: [Testing in Go](https://golang.org/pkg/testing/).