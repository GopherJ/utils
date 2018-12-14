package util

/**
 * atoi.go
 * 2017-12-27 14:30:27
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "strconv"
)

func MustAtoi(s string)(int) {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}
