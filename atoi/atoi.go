package gocontrib

import "strconv"

func AtoiImproved(s string) (int, error) {
	const fnAtoi = "Atoi"

	sLen := len(s)

	if sLen > 0 {
		// Could be a fast path for small integers that fit int type.
		first := 0
		switch s[0] {
		// Check 1st char for sign.
		case '+', '-':
			if sLen == 1 {
				return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
			}
			first = 1
		}

		if intSize == 32 && sLen < 10 + first || intSize == 64 && sLen < 19 + first {
			// Fast path for small integers that fit int type.
			n := 0
			for _, ch := range s[first:] {
				ch -= '0'
				if ch > 9 {
					return 0, &strconv.NumError{fnAtoi, s, strconv.ErrSyntax}
				}
				n = n*10 + int(ch)
			}
			if s[0] == '-' {
				n = -n
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
