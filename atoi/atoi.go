package gocontrib

import "strconv"

const intSize = 32 << (^uint(0) >> 63)

func AtoiImproved(s string) (int, error) {
	const fnAtoi = "Atoi"

	sLen := len(s)

	if intSize == 32 && 0 < sLen && sLen < 10 || intSize == 64 && 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		first := 0
		neg := s[0] == '-'
		if neg || s[0] == '+' {
			// Check 1st char for sign.
			// case '+', '-':
			if sLen < 2 {
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			first = 1
		}
		n := 0
		for i := first; i < sLen; i++ {
			ch := s[i] - '0'
			if ch > 9 {
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			n = n*10 + int(ch)
		}
		if neg {
			n = -n
		}
		return n, nil
	}

	// Slow path for invalid or big integers.
	i64, err := strconv.ParseInt(s, 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtoi
	}
	return int(i64), err
}
