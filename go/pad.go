package util

/**
 * pad.go
 * 2017-12-27 13:04:36
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "strings"
    "unicode/utf8"
    "math"
)

func PL(s, r string, n int) string {
    L := utf8.RuneCountInString(s)
    if n > L {
        return s + strings.Repeat(r, n - L)
    } else {
        return s
    }
}

func PR(s, r string, n int) string {
    L := utf8.RuneCountInString(s)
    if n > L {
        return strings.Repeat(r, n - L) + s
    } else {
        return s
    }
}

func PC(s, r string, n int) string {
    L := utf8.RuneCountInString(s)
    if n > L {
        le := int(math.Floor(float64(n - L) / 2))
        ri := int(math.Ceil(float64(n - L) / 2))
        return strings.Repeat(r, le) + s + strings.Repeat(r, ri)
    } else {
        return s
    }
}
