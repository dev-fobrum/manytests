package tasks

import (
	"errors"
	"fmt"
	"strconv"
)

func convertTimeAMPMToMilitary(s string) (string, error) {
	sl := len(s)

	if sl != 10 {
		return "", errors.New(fmt.Sprintf("Expect a 10 lenght string and recevied %d", sl))
	}

	hour := s[0:2]
	format := s[sl-2 : sl]
	rest := s[2 : sl-2]

	if format == "PM" {
		res, err := strconv.Atoi(hour)
		if err != nil {
			return "", errors.New("Error converting hour to int")
		}

		if res == 12 {
			hour = "12"
		} else {
			hour = strconv.Itoa(res + 12)
		}

		convertedTime := hour + rest

		return convertedTime, nil
	}

	if hour == "12" {
		hour = "00"
	}

	convertedTime := hour + rest

	return convertedTime, nil
}
