package myatoi

import (
	"math"
)

func myAtoi(s string) int {
	var res int64
	sign, letter, number := false, false, false
	var flag int64 = 1
	for _, v := range s {
		if v == ' ' {
			if sign || letter || number {
				break
			}
		} else if v == '-' || v == '+' {
			if sign || letter || number {
				break
			}
			sign = true
			if v == '-' {
				flag = -1
			}
		} else if v >= '0' && v <= '9' && !letter {
			number = true
			res = res*10 + (int64(v) - int64('0'))
		} else {
			letter = true
		}
		if res*flag < math.MinInt32 {
			return math.MinInt32
		}
		if res*flag > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return int(res * flag)
}
