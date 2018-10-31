package gocontrib

import "strconv"

const intSize = 32 << (^uint(0) >> 63)

// Is used by Atoi() and may be used in future by ParseInt()
// to determine if fast path is available for small integers that fit int type.
func isFastPathAtoi(sLen int) bool {
	return intSize == 32 && 0 < sLen && sLen < 10 || intSize == 64 && 0 < sLen && sLen < 19
}

func AtoiImproved(s string) (int, error) {
	const fnAtoi = "Atoi"

	sLen := len(s)

	if isFastPathAtoi(sLen) {
		// Fast path for small integers that fit int type.
		startPos := 0
		neg := s[0] == '-'
		if neg || s[0] == '+' {
			if sLen < 2 {
				// Sign only strings are bad.
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			startPos = 1
		}
		n := int(s[startPos] - '0') // First time no need to multiply previous n by 10.
		if n > 9 {
			return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
		}
		for i := startPos+1; i < sLen; i++ {
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
