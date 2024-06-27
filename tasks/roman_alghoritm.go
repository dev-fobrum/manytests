package tasks

import (
	"errors"
	"fmt"
	"regexp"
)

func isValidRoman(s string) bool {
	romanRegex := `^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`

	re := regexp.MustCompile(romanRegex)

	return re.MatchString(s)
}

func executeRomanAlghoritm(s string) (int, error) {
	if !isValidRoman(s) {
		return 0, errors.New("Is not a valid number")
	}

	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i < n-1 && romanValues[s[i]] < romanValues[s[i+1]] {
			total -= romanValues[s[i]]
		} else {
			total += romanValues[s[i]]
		}
	}

	return total, nil
}

func RunRomanAlghoritm() {
	total, err := executeRomanAlghoritm("IIII")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(total)

}
