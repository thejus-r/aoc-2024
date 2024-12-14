package utils

import "strconv"

// This function returns absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func SimpleMod(a, b int) int {
	r := a % b
	if r < 0 && b > 0 {
		r += b
	}
	return r
}
