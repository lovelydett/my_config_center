package utils

import "strconv"

func Atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic("Unable to convert string to int")
	}

	return res
}
