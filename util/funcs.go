package util

import "strconv"

func ToF64(s string) float64 {
	f64, _ := strconv.ParseFloat(s, 64)
	return f64
}

func ToF64P(s *string) float64 {
	if s == nil {
		return 0
	}
	return ToF64(*s)
}
