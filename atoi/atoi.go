package gocontrib

import "strconv"

func AtoiImproved(s string) (int, error) {
	const fnAtoi = "Atoi"

	sLen := len(s)

	if strconv.IntSize == 32 && (0 < sLen && sLen < 10) ||
		strconv.IntSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		switch s[0] {
		case '-':
			if sLen == 1 {
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			n := 0
			for _, ch := range []byte(s[1:]) {
				ch -= '0'
				if ch > 9 {
					return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
				}
				n = n*10 + int(ch)
			}
			return -n, nil
		case '+':
			if sLen == 1 {
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			n := 0
			for _, ch := range []byte(s[1:]) {
				ch -= '0'
				if ch > 9 {
					return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
				}
				n = n*10 + int(ch)
			}
			return n, nil
		default:
			n := 0
			for _, ch := range []byte(s) {
				ch -= '0'
				if ch > 9 {
					return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
				}
				n = n*10 + int(ch)
			}
			return n, nil
		}
	}

	// Slow path for invalid or big integers.
	i64, err := strconv.ParseInt(s, 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtoi
	}
	return int(i64), err
}
