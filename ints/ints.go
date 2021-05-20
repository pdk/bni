package ints

import "strconv"

// Max returns the largest int. Returns 0 for empty slice.
func Max(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}

	return m
}

// Min returns the smallest integer. Returns 0 for empty slice.
func Min(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}

	return m
}

// Sum returns the total of adding all the integers. Returns 0 for empty slice.
func Sum(vals []int) int {

	s := 0
	for _, v := range vals {
		s += v
	}

	return s
}

// Product returns the result of multiplying all the integers. Returns 1 for empty slice.
func Product(vals []int) int {

	s := 1
	for _, v := range vals {
		s *= v
	}

	return s
}

// Strings returns the ints as strings.
func Strings(vals []int) []string {

	s := make([]string, 0, len(vals))

	for _, v := range vals {
		s = append(s, strconv.Itoa(v))
	}

	return s
}

// Parse converts a slice of strings to a slice of ints.
func Parse(vals []string) ([]int, error) {

	ints := make([]int, 0, len(vals))

	for _, s := range vals {
		i, err := strconv.Atoi(s)
		if err != nil {
			return ints, err
		}
		ints = append(ints, i)
	}

	return ints, nil
}
