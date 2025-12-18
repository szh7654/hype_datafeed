package util

import (
	"strconv"
	"time"
)

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

func MaxTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}
